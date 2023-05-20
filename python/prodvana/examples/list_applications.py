import argparse

from prodvana.client import Client, make_channel
from prodvana.proto.prodvana.application.application_manager_pb2 import (
    ListApplicationsReq,
)


def main() -> None:
    ap = argparse.ArgumentParser()
    ap.add_argument(
        "--api-token", help="Prodvana API token. if not passed, defaults to PVN_TOKEN."
    )
    ap.add_argument(
        "--org",
        required=True,
        help="Organization slug for your instance of Prodvana. This is the part of your URL before .prodvana.io.",
    )
    args = ap.parse_args()
    with make_channel(org=args.org, api_token=args.api_token) as channel:
        client = Client(channel=channel)
        resp = client.application_manager.ListApplications(ListApplicationsReq())
        for app in resp.applications:
            print(app.meta.name)


if __name__ == "__main__":
    main()
