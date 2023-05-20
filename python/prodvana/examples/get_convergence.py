import argparse
from typing import NamedTuple

from prodvana.client import Client, make_channel
from prodvana.proto.prodvana.desired_state.manager_pb2 import (
    GetServiceDesiredStateConvergenceSummaryReq,
)
from prodvana.proto.prodvana.desired_state.model.desired_state_pb2 import Status, Type


class HashableIdentifier(NamedTuple):
    type: "Type.V"
    name: str


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
    ap.add_argument("app", help="Application that contains the service")
    ap.add_argument("service", help="Service to get convergence status for")
    args = ap.parse_args()
    app = args.app
    service = args.service
    with make_channel(org=args.org, api_token=args.api_token) as channel:
        client = Client(channel=channel)
        resp = client.desired_state_manager.GetServiceDesiredStateConvergenceSummary(
            GetServiceDesiredStateConvergenceSummaryReq(
                application=app, service=service
            )
        )
        graph = {
            HashableIdentifier(type=entity.id.type, name=entity.id.name): entity
            for entity in resp.summary.entity_graph.entities
        }
        svc_entity = graph[
            HashableIdentifier(
                type=resp.summary.entity_graph.root.type,
                name=resp.summary.entity_graph.root.name,
            )
        ]
        print(f"status: {Status.Name(svc_entity.status)}")
        for child in svc_entity.dependencies:
            if child.type != Type.SERVICE_INSTANCE:
                continue
            child_entity = graph[HashableIdentifier(type=child.type, name=child.name)]
            print(
                f"release channel: {child_entity.desired_state.service_instance.release_channel}, status: {Status.Name(child_entity.status)}"
            )


if __name__ == "__main__":
    main()
