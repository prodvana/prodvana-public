import argparse
import json
import logging
import os
from dataclasses import dataclass
from datetime import datetime, timedelta
from typing import Dict, List, Optional

import pytz
import requests
from dxf import DXF  # type: ignore
from github import Commit, Github, PullRequest
from google.cloud import spanner  # type: ignore

PRODVANA_ORG_ID = "org-pvn-prime"

# prodvana services as of 2022-11-23
IN_SCOPE_SERVICE_NAMES = [
    "agent-interaction",
    "apiserver",
    "cache-warmer",
    "insights-generator",
    "k8s-poller",
    "ui",
]
# production release channels as of 2022-11-23
IN_SCOPE_RELEASE_CHANNELS = [
    "org-02bc46d297454948957488ec4d2fd132",
    "org-0efd6ace5fa947f3b4a1fdd2f6baec70",
    "org-f12f30336b914efb9603a460f8412dea",
    "org-fc6ce3f1ee564f3780c6e65bc933f39a",
    "org-5684ff5b04e74e588298ed592cc3d41d",
]

PRODVANA_IMAGE_PREFIX = "us-central1-docker.pkg.dev/pvn-infra/pvn/"

# github api has very aggressive rate limits
GITHUB_API_RETRY_LIMIT = 10

UTC = pytz.UTC


@dataclass
class DeployDetails:
    service: str
    release_channel: str
    repo: str
    docker_tag: str
    deploy_start_ts: datetime

    docker_digest: Optional[str] = None
    git_hash: Optional[str] = None

    def key(self) -> str:
        return f"{self.repo}:{self.docker_tag}:{self.deploy_start_ts}"


@dataclass
class Finding:
    git_hash: str
    pr: PullRequest.PullRequest
    review_ts: Optional[datetime] = None


@dataclass
class Findings:
    findings: List[Finding]
    deploy_start_ts: datetime
    root_git_hash: str


# Returns None if PR was not reviewed, otherwise returns the timestamp of when it was reviewed
def pull_reviewed_at(
    commit: Commit.Commit, pull: PullRequest.PullRequest
) -> Optional[datetime]:
    for issue_comment in pull.get_issue_comments():
        if issue_comment.user == commit.author:
            continue
        if issue_comment.body.strip().lower() == "lgtm":
            return issue_comment.created_at
    for review in pull.get_reviews():
        if review.user == commit.author:
            continue
        if review.state == "APPROVED" or review.body.strip().lower() == "lgtm":
            return review.submitted_at
    for review_comment in pull.get_review_comments():
        if review_comment.user == commit.author:
            continue
        if review_comment.body.strip().lower() == "lgtm":
            return review_comment.created_at
    return None


def main() -> None:
    logging.basicConfig(level=logging.INFO)
    ap = argparse.ArgumentParser()
    ap.add_argument("--repo", default="prodvana/prodvana")
    ap.add_argument(
        "--gh-access-token",
        help="Github access token, defaults to value of env variable GITHUB_TOKEN",
    )
    ap.add_argument(
        "--docker-username",
        help="Docker username, defaults to DXF_USERNAME",
    )
    ap.add_argument(
        "--docker-password",
        help="Docker password, defaults to DXF_PASSWORD",
    )
    ap.add_argument(
        "--start", help="range start date in YYYY-mm-dd format", required=True
    )
    ap.add_argument("--end", help="range end date in YYYY-mm-dd format", required=True)
    ap.add_argument(
        "--instance-id", help="Spanner instance id", default="pvn-spanner-prime-f825071"
    )
    ap.add_argument("--database-id", help="Spanner database id", default="prodvana-db")
    args = ap.parse_args()

    def auth(dxf_cli: DXF, response: requests.Response) -> None:
        dxf_cli.authenticate(
            username=args.docker_username or os.getenv("DXF_USERNAME"),
            password=args.docker_password or os.getenv("DXF_PASSWORD"),
            response=response,
        )

    cli = spanner.Client(project="pvn-infra")
    db = cli.instance(args.instance_id).database(args.database_id)

    # fetch service ids for services
    service_to_ids = {}
    ids_to_service = {}
    with db.snapshot() as snapshot:
        results = snapshot.execute_sql(
            f"""
            SELECT Name, ServiceId
            FROM Services
            WHERE OrgId = @orgId
            AND name IN ({",".join([f"'{svc}'" for svc in IN_SCOPE_SERVICE_NAMES])});
            """,
            params={"orgId": PRODVANA_ORG_ID},
            param_types={
                "orgId": spanner.param_types.STRING,
            },
        )

        for (name, id) in results:
            service_to_ids[name] = id
            ids_to_service[id] = name

    # look up release channel ids for Production RCs, since only Production deploys are in scope
    prod_release_channel_ids = []
    with db.snapshot() as snapshot:
        results = snapshot.execute_sql(
            f"""
            SELECT ReleaseChannelId
            FROM ReleaseChannels
            WHERE OrgId = @orgId
            AND name IN ({",".join([f"'{rc}'" for rc in IN_SCOPE_RELEASE_CHANNELS])});
            """,
            params={"orgId": PRODVANA_ORG_ID},
            param_types={
                "orgId": spanner.param_types.STRING,
            },
        )

        for row in results:
            prod_release_channel_ids.append(row[0])

    repotag_to_details: Dict[str, DeployDetails] = {}
    repos_to_tag_scrape = set()
    with db.snapshot() as snapshot:
        # fetch all pushes with timestamps and configs
        push_infos = snapshot.execute_sql(
            f"""
            SELECT
                sip.ServiceID,
                sv.ServiceVersion,
                ANY_VALUE(sv.ConfigJson),
                MAX(sip.StartTimestamp)
            FROM ServiceInstancePushes sip
            JOIN ServiceVersions sv ON sip.ServiceVersion = sv.ServiceVersion
            JOIN ReleaseChannels rc ON sip.ReleaseChannelId = rc.ReleaseChannelId
            WHERE
                sip.OrgId = @orgId AND
                sip.ServiceID IN ({','.join([f"'{id}'" for id in service_to_ids.values()])}) AND
                sip.ReleaseChannelId IN ({','.join([f"'{id}'" for id in prod_release_channel_ids])}) AND
                sip.State="succeeded" AND
                TIMESTAMP_DIFF(sip.StartTimestamp, TIMESTAMP '{args.start}', DAY) > 0 AND
                TIMESTAMP_DIFF(sip.StartTimestamp, TIMESTAMP '{args.end}', DAY) <= 0
            GROUP BY sip.ServiceID, sv.ServiceVersion;
        """,
            params={"orgId": PRODVANA_ORG_ID},
            param_types={
                "orgId": spanner.param_types.STRING,
            },
        )

        registry_cache: Dict[str, DXF] = {}
        # parse the configs to get docker image urls and parse out the git hashes from the tag
        # if the url is an image digest look up the tag
        for (sid, sv, config, ts) in push_infos:
            image_url = config["programs"][0]["image"].replace(
                PRODVANA_IMAGE_PREFIX, ""
            )
            digest = None
            git_hash = None

            # image digest "tags" will look like @sha256:<hash>
            digest_parts = image_url.split("@")
            if len(digest_parts) == 2:
                [repo, tag] = digest_parts
                repos_to_tag_scrape.add(repo)
                if repo in registry_cache:
                    dxf = registry_cache[repo]
                else:
                    dxf = DXF(
                        "us-central1-docker.pkg.dev",
                        repo=f"pvn-infra/pvn/{repo}",
                        auth=auth,
                    )
                    registry_cache[repo] = dxf
                manifest = json.loads(dxf.get_manifest(tag))
                digest = manifest["config"]["digest"] or tag
            else:
                [repo, tag] = image_url.split(":")
                git_hash = tag

            details = DeployDetails(
                service=ids_to_service[sid],
                release_channel="",
                repo=repo,
                docker_tag=tag,
                deploy_start_ts=ts,
                docker_digest=digest,
                git_hash=git_hash,
            )
            repotag_to_details[details.key()] = details

    # now lets look for all the image digests (@sha256:<hash>) and resolve them to the actual tags.
    # the docker registry api doesn't have an easy way to reverse this without fetching all the tags
    # to figure out what digest they map to.
    digest_to_tags_by_repo = {}
    for repo in repos_to_tag_scrape:
        if repo in registry_cache:
            dxf = registry_cache[repo]
        else:
            dxf = DXF(
                "us-central1-docker.pkg.dev",
                repo=f"pvn-infra/pvn/{repo}",
                auth=auth,
            )
            registry_cache[repo] = dxf
        digest_to_tags = {}
        all_tags = dxf.list_aliases()
        logging.info(f"getting repo:{repo} tag info tag count {len(all_tags)}")
        for tag in all_tags:
            if tag == "latest":
                # skip latest since it's never used for deployment
                continue
            manifest = json.loads(dxf.get_manifest(tag))
            digest = manifest["config"]["digest"]
            if digest in digest_to_tags:
                logging.warning(f"found another tag:{tag} for digest:{digest}")
            else:
                digest_to_tags[digest] = tag
        digest_to_tags_by_repo[repo] = digest_to_tags

    gh = Github(
        args.gh_access_token or os.getenv("GITHUB_TOKEN"), retry=GITHUB_API_RETRY_LIMIT
    )
    repo = gh.get_repo(args.repo)

    # sort details on deploy time so we can keep track of the previous push version
    sorted_details = sorted(
        repotag_to_details.values(), key=lambda d: d.deploy_start_ts
    )
    last_commit_by_repo: Dict[str, DeployDetails] = {}
    not_reviewed_by_repohash: Dict[str, Findings] = {}
    for details in sorted_details:
        # resolve digest to tag if needed
        if details.docker_digest:
            if (
                details.repo not in digest_to_tags_by_repo
                or details.docker_digest not in digest_to_tags_by_repo[details.repo]
            ):
                logging.warning(
                    f"[skipping] could not find tag for digest: ({details.repo}, {details.docker_digest})"
                )
                continue
            details.git_hash = digest_to_tags_by_repo[details.repo][
                details.docker_digest
            ]
        if details.repo in last_commit_by_repo:
            last_hash = last_commit_by_repo[details.repo].git_hash
            commits = repo.compare(last_hash, details.git_hash).commits
        else:
            commit = repo.get_commit(sha=details.git_hash)
            commits = [commit]

        last_commit_by_repo[details.repo] = details

        findings = []
        for commit in commits:
            pulls = commit.get_pulls()
            for pull in pulls:
                review_ts = pull_reviewed_at(commit, pull)
                if review_ts is None:
                    findings.append(
                        Finding(
                            git_hash=commit.sha,
                            pr=pull,
                        )
                    )
                else:
                    review_ts = UTC.localize(review_ts)
                    deploy_ts_plus_24h = details.deploy_start_ts + timedelta(hours=24)
                    if deploy_ts_plus_24h < review_ts:
                        findings.append(
                            Finding(
                                git_hash=commit.sha,
                                pr=pull,
                                review_ts=review_ts,
                            )
                        )
        if len(findings) > 0:
            git_hash = details.git_hash or "<missing>"
            key = f"{details.repo}:{details.git_hash}"
            not_reviewed_by_repohash[key] = Findings(
                findings=findings,
                deploy_start_ts=details.deploy_start_ts,
                root_git_hash=git_hash,
            )

    for rh, fds in not_reviewed_by_repohash.items():
        print(f"{rh} {fds.root_git_hash} {fds.deploy_start_ts}")
        for fd in fds.findings:
            print(
                f"\t{fd.git_hash} https://github.com/prodvana/prodvana/pull/{fd.pr.number} {fd.review_ts}"
            )
        print("------------------")


if __name__ == "__main__":
    main()
