import argparse
import logging
import os
from typing import Any, NamedTuple, cast

from github import Github, Repository
from google.cloud import spanner  # type: ignore
from prodvana_sdk.client import Client, make_channel
from prodvana_sdk.utils import desired_states as desired_states_util
from prodvana_sdk.utils import service_config as service_config_util

# github api has very aggressive rate limits
GITHUB_API_RETRY_LIMIT = 10
PRODVANA_REGISTRY_NAME = "gcp-us-central1"
PRODVANA_REGISTRY_ID = "int-97e2fc2bbb894ed993b1cf71f4bdcfe4"
PRODVANA_IMAGE_PREFIX = "us-central1-docker.pkg.dev/pvn-infra"
MIGRATIONS_DIR = "go/db/dbtestutil/migrations"


class MigrationNotApplied(Exception):
    pass


class SpannerResource(NamedTuple):
    project: str
    instance: str
    database: str


def parse_spanner_resource(resource: str) -> SpannerResource:
    parts = resource.strip("/").split("/")
    assert len(parts) == 6
    assert parts[0] == "projects"
    assert parts[2] == "instances"
    assert parts[4] == "databases"
    project = parts[1]
    instance = parts[3]
    database = parts[5]
    return SpannerResource(project=project, instance=instance, database=database)


def get_latest_schema_version(spanner_snapshot: Any) -> int:
    results = spanner_snapshot.read(
        table="SchemaMigrations",
        columns=("Version", "Dirty"),
        keyset=spanner.KeySet(all_=True),
    )
    rows = list(results)
    assert len(rows) == 1, f"Unexpected number of rows in SchemaMigrations: {rows}"
    version, dirty = rows[0]
    assert not dirty, "Migration state is dirty, please fix."
    return cast(int, version)


def assert_migration_applied_for_commit(
    latest_schema_version: int, repo: Repository.Repository, commit: str
) -> None:
    contents = repo.get_contents(path=MIGRATIONS_DIR, ref=commit)
    assert isinstance(contents, list)
    latest_version_on_commit = None
    for content in contents:
        base, ext = os.path.splitext(content.name)
        if ext != ".sql":
            continue
        base_int = int(base)
        if latest_version_on_commit is None:
            latest_version_on_commit = base_int
        else:
            latest_version_on_commit = max(base_int, latest_version_on_commit)
    assert (
        latest_version_on_commit is not None
    ), f"Did not find any database migrations on {commit}"
    if latest_schema_version < latest_version_on_commit:
        raise MigrationNotApplied(
            f"Migration not run for commit {commit}. Latest schema version: {latest_schema_version}, required schema version: {latest_version_on_commit}"
        )


class SkipException(Exception):
    """
    Exceptions in commit extraction that should result in a skip instead of a hard failure
    """

    pass


class NotProdvanaCodeException(SkipException):
    pass


class NoTagsException(SkipException):
    pass


def get_commit_from_image(
    details: service_config_util.ImageDetails,
) -> str:

    if details.image_url is not None:
        image = details.image_url
        if not image.startswith(PRODVANA_IMAGE_PREFIX):
            raise NotProdvanaCodeException(
                f"Image {image} does not contain prodvana code"
            )
        if "@sha256:" in image:
            # Prodvana used to enforce digest, so this is here to ease the transition.
            raise NoTagsException(f"Image {image} uses a digest instead of a tag")
        _, tag = image.rsplit(":", 1)
        assert (
            len(tag) == 40
        ), f"image tag for {image} does not look like it is a git sha"
        return tag
    elif details.tag is not None:
        tag_details = details.tag
        assert (
            tag_details.registry_name or tag_details.registry_id
        ), f"Missing container registry information for image with details {details}"

        if (
            tag_details.registry_name
            and tag_details.registry_name != PRODVANA_REGISTRY_NAME
        ) or (
            tag_details.registry_id and tag_details.registry_id != PRODVANA_REGISTRY_ID
        ):
            raise NotProdvanaCodeException(
                f"Image tag {tag_details.tag} does not contain prodvana code, from {details}"
            )
        return tag_details.tag
    else:
        raise NoTagsException(f"No image information found from details: {details}")


def main() -> None:
    logging.basicConfig(level=logging.INFO)
    ap = argparse.ArgumentParser()
    ap.add_argument("--repo", default="prodvana/prodvana")
    ap.add_argument(
        "--access-token",
        help="Access token, defaults to value of env variable GITHUB_TOKEN",
    )
    ap.add_argument(
        "--spanner-resource",
        help="Spanner resource path, defaults to value of env variable PVN_SPANNER_RESOURCE",
    )
    ag = ap.add_mutually_exclusive_group(required=True)
    ag.add_argument("--commit", nargs="+", help="Commit to check up to")
    ag.add_argument(
        "--commits-from-prodvana",
        action="store_true",
        help="Derive commits from prodvana config's docker images. Assumes this is being run in the context of a Prodvana custom pipeline task.",
    )
    args = ap.parse_args()
    gh = Github(
        args.access_token or os.getenv("GITHUB_TOKEN"), retry=GITHUB_API_RETRY_LIMIT
    )
    spanner_resource_str = args.spanner_resource or os.getenv("PVN_SPANNER_RESOURCE")
    assert (
        spanner_resource_str
    ), "Must pass --spanner-resource or set PVN_SPANNER_RESOURCE"
    spanner_resource = parse_spanner_resource(spanner_resource_str)
    spanner_client = spanner.Client(project=spanner_resource.project)
    spanner_instance = spanner_client.instance(spanner_resource.instance)
    spanner_db = spanner_instance.database(spanner_resource.database)
    target_commits = set[str]()
    if args.commits_from_prodvana:
        channel = make_channel()
        with channel:
            client = Client(channel)
            desired_state_summary = desired_states_util.get_desired_state_details(
                client
            )
            target_images = service_config_util.get_target_images_for_desired_state(
                client, desired_state_summary
            )

            for image in target_images:
                commit_from_image = get_commit_from_image(image)
                target_commits.add(commit_from_image)
    else:
        target_commits.update(args.commit)
    assert target_commits
    repo = gh.get_repo(args.repo)
    with spanner_db.snapshot() as snapshot:
        latest_schema_version = get_latest_schema_version(snapshot)

    for commit in target_commits:
        logging.info(
            f"Checking if {commit} requires a newer migration version than {latest_schema_version}"
        )
        assert_migration_applied_for_commit(
            latest_schema_version=latest_schema_version, repo=repo, commit=commit
        )

    logging.info("All required migrations are already applied.")


if __name__ == "__main__":
    main()
