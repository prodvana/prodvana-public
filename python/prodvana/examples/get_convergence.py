import argparse
from typing import Iterator, Mapping, NamedTuple, Optional

from prodvana.client import Client, make_channel
from prodvana.proto.prodvana.desired_state.manager_pb2 import (
    GetServiceDesiredStateConvergenceSummaryReq,
)
from prodvana.proto.prodvana.desired_state.model.desired_state_pb2 import (
    SignalType,
    Status,
    Type,
)
from prodvana.proto.prodvana.desired_state.model.entity_pb2 import Entity


class HashableIdentifier(NamedTuple):
    type: "Type.V"
    name: str


class MissingApproval(NamedTuple):
    topic: str
    signal_type: "SignalType.V"
    desired_state_id: str


def find_missing_approval(
    graph: Mapping[HashableIdentifier, Entity], release_channel: HashableIdentifier
) -> Optional[MissingApproval]:
    release_channel_entity = graph[release_channel]
    for dep in release_channel_entity.dependencies:
        if dep.type == Type.MANUAL_APPROVAL:
            manual_approval_entity = graph[
                HashableIdentifier(type=dep.type, name=dep.name)
            ]
            if manual_approval_entity.status == Status.CONVERGING:
                return MissingApproval(
                    topic=release_channel_entity.desired_state.service_instance.release_channel,
                    signal_type=SignalType.SIGNAL_MANUAL_APPROVAL,
                    desired_state_id=release_channel_entity.root_desired_state_id,
                )

    # No manual approval entities.
    # Check if anything in the tree has missing approval.

    def visit(
        graph: Mapping[HashableIdentifier, Entity], node: HashableIdentifier
    ) -> Iterator[HashableIdentifier]:
        for dep in graph[node].dependencies:
            yield HashableIdentifier(type=dep.type, name=dep.name)

    for node in visit(graph, release_channel):
        missing_approval = graph[node].missing_approval
        if missing_approval is not None and len(missing_approval.topic) > 0:
            return MissingApproval(
                topic=missing_approval.topic,
                signal_type=missing_approval.signal_type,
                desired_state_id=release_channel_entity.root_desired_state_id,
            )

    return None


def main() -> None:
    ap = argparse.ArgumentParser()
    ap.add_argument(
        "--api-token", help="Prodvana API token. if not passed, defaults to PVN_TOKEN."
    )
    ap.add_argument(
        "--org",
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

            child_id = HashableIdentifier(type=child.type, name=child.name)
            child_entity = graph[child_id]
            print(
                f"release channel: {child_entity.desired_state.service_instance.release_channel}, status: {Status.Name(child_entity.status)}"
            )

            if child_entity.status == Status.WAITING_MANUAL_APPROVAL:
                print(
                    f"Missing approval for {child_entity.desired_state.service_instance.release_channel}:"
                )
                missing_approval = find_missing_approval(graph, child_id)
                if missing_approval:
                    print(
                        f"\tpvnctl services approve --app {app} {service} {missing_approval.topic} {missing_approval.desired_state_id} --signal-type {SignalType.Name(missing_approval.signal_type)}"
                    )


if __name__ == "__main__":
    main()
