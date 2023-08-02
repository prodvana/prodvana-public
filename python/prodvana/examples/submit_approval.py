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
        help="Organization slug for your instance of Prodvana. This is the part of your URL before .prodvana.io.",
    )
    ap.add_argument("ds_id", help="Desired state id to approve for")
    ap.add_argument(
        "topic", help="Topic to approve (release channel name or missing approval id)"
    )
    ap.add_argument(
        "signal_type", default="SIGNAL_MANUAL_APPROVAL", help="Signal type to approve"
    )
    args = ap.parse_args()
    ds_id = args.ds_id
    topic = args.topic
    signal_type = args.signal_type
    with make_channel(org=args.org, api_token=args.api_token) as channel:
        client = Client(channel=channel)

        client.desired_state_manager.SetManualApproval(
            SetManualApprovalReq(
                desired_state_id=ds_id,
                topic=topic,
                signal_type=signal_type,
                # uncomment to reject instead of approve
                # reject=True,
            )
        )


if __name__ == "__main__":
    main()
