import argparse

from prodvana.client import Client, make_channel
from prodvana.proto.prodvana.desired_state.manager_pb2 import SetManualApprovalReq


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
    ap.add_argument("ds_id", help="Desired state id to approve for")
    ap.add_argument("release_channel", help="Release channel name to approve")
    args = ap.parse_args()
    ds_id = args.ds_id
    release_channel = args.release_channel
    with make_channel(org=args.org, api_token=args.api_token) as channel:
        client = Client(channel=channel)

        client.desired_state_manager.SetManualApproval(
            SetManualApprovalReq(
                desired_state_id=ds_id,
                topic=release_channel,
                # uncomment to reject instead of approve
                # reject=True,
            )
        )


if __name__ == "__main__":
    main()
