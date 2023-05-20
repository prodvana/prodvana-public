import argparse

from prodvana.client import Client, make_channel
from prodvana.proto.prodvana.service.service_manager_pb2 import ListServicesReq


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
    ap.add_argument("app", help="Application to list services from.")
    args = ap.parse_args()
    with make_channel(org=args.org, api_token=args.api_token) as channel:
        client = Client(channel=channel)
        resp = client.service_manager.ListServices(
            ListServicesReq(application=args.app)
        )
        for svc in resp.services:
            print(svc.meta.name)


if __name__ == "__main__":
    main()
