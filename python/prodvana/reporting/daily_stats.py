import argparse
import datetime
import json
import logging
import os
from typing import NamedTuple

from google.cloud import spanner  # type: ignore
from slack_sdk import WebClient


class OrganizationStats(object):
    org_name: str
    org_id: str
    daily_successes: int
    daily_failures: int
    active: bool

    def __init__(  # type: ignore
        self,
        org_id: str,
        org_name: str,
        daily_successes=0,
        daily_failures=0,
        active=False,
    ) -> None:
        self.org_id = org_id
        self.org_name = org_name
        self.daily_successes = daily_successes
        self.daily_failures = daily_failures
        self.active = active
        return None

    def __str__(self) -> str:
        return "%s %s %d %d %s" % (
            self.org_id,
            self.org_name,
            self.daily_successes,
            self.daily_failures,
            self.active,
        )


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


def main() -> None:

    logging.basicConfig(level=logging.INFO)
    ap = argparse.ArgumentParser()
    ap.add_argument(
        "--spanner-resource",
        help="Spanner resource path, defaults to value of env variable PVN_SPANNER_RESOURCE",
    )
    ap.add_argument(
        "--stats_slack_channel",
        help="Where to spam daily stats",
        default="prodvana-stats",
    )
    args = ap.parse_args()

    SLACK_TOKEN = os.environ.get("PVN_SLACK_STATSBOT_TOKEN")
    assert SLACK_TOKEN, "Must provide slack token"
    slack_channel = args.stats_slack_channel

    spanner_resource_str = args.spanner_resource or os.getenv("PVN_SPANNER_RESOURCE")
    assert (
        spanner_resource_str
    ), "Must pass --spanner-resource or set PVN_SPANNER_RESOURCE"
    spanner_resource = parse_spanner_resource(spanner_resource_str)
    spanner_client = spanner.Client(project=spanner_resource.project)
    spanner_instance = spanner_client.instance(spanner_resource.instance)
    spanner_db = spanner_instance.database(spanner_resource.database)

    orgs = {}
    with spanner_db.snapshot() as snapshot:
        results = snapshot.execute_sql("SELECT OrgID, Name from Organizations;")
    for row in results:
        orgs[row[0]] = OrganizationStats(row[0], row[1])

    today = datetime.date.today()
    yesterday = today - datetime.timedelta(days=1)
    today_ts = datetime.datetime.strptime(str(today), "%Y-%m-%d").isoformat() + "Z"
    yesterday_ts = (
        datetime.datetime.strptime(str(yesterday), "%Y-%m-%d").isoformat() + "Z"
    )

    # Get Successes
    with spanner_db.snapshot() as snapshot:
        results = snapshot.execute_sql(
            "SELECT OrgId, count(*) as deploys FROM desiredstates WHERE status = 2 AND EntityType = 2 and CreationTimestamp > @yesterday AND CreationTimestamp < @today Group BY OrgId",
            params={"yesterday": yesterday_ts, "today": today_ts},
            param_types={
                "yesterday": spanner.param_types.STRING,
                "today": spanner.param_types.STRING,
            },
        )

    for row in results:
        if row[0] in orgs.keys():
            orgs[row[0]].daily_successes = row[1]
            orgs[row[0]].active = True

    # Get Failures
    with spanner_db.snapshot() as snapshot:
        results = snapshot.execute_sql(
            "SELECT OrgId, count(*) as deploys FROM desiredstates WHERE status = 3 AND EntityType = 2 and CreationTimestamp > @yesterday AND CreationTimestamp < @today Group BY OrgId",
            params={"yesterday": yesterday_ts, "today": today_ts},
            param_types={
                "yesterday": spanner.param_types.STRING,
                "today": spanner.param_types.STRING,
            },
        )

    for row in results:
        if row[0] in orgs.keys():
            orgs[row[0]].daily_failures = row[1]
            orgs[row[0]].active = True

    slack_client = WebClient(SLACK_TOKEN)
    slack_messages = []
    message = {
        "type": "header",
        "text": {
            "type": "plain_text",
            "text": "Prodvana Daily Stats for %s - %s" % (yesterday, today),
        },
    }
    slack_messages.append(message)
    for org_name in orgs:
        if orgs[org_name].active:
            message = {
                "type": "section",
                "text": {
                    "type": "plain_text",
                    "text": "%s: %s - %s"
                    % (
                        orgs[org_name].org_name,
                        orgs[org_name].daily_successes,
                        orgs[org_name].daily_failures,
                    ),
                },
            }
            slack_messages.append(message)
    message_blocks = json.dumps(slack_messages)

    slack_client.chat_postMessage(
        channel="#" + slack_channel, blocks=message_blocks, text="Prodvana Stats"
    )


if __name__ == "__main__":
    main()
