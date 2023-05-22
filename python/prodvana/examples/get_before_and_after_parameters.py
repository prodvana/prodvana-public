import argparse

from prodvana.client import Client, make_channel
from prodvana.proto.prodvana.common_config.parameters_pb2 import ParameterValue
from prodvana.proto.prodvana.desired_state.manager_pb2 import (
    GetDesiredStateConvergenceReq,
)
from prodvana.proto.prodvana.service.service_manager_pb2 import GetMaterializedConfigReq


def print_param_value(param: ParameterValue) -> None:
    print(f"parameter {param.name}: ", end="")
    if param.string:
        print(param.string)
    elif param.int:
        print(f"{param.int}")
    elif param.docker_image_tag:
        print(param.docker_image_tag)
    else:
        raise Exception(f"unrecognized parameter: {param}")


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
    args = ap.parse_args()
    ds_id = args.ds_id
    with make_channel(org=args.org, api_token=args.api_token) as channel:
        client = Client(channel=channel)

        resp = client.desired_state_manager.GetDesiredStateConvergenceSummary(
            GetDesiredStateConvergenceReq(
                desired_state_id=ds_id,
            )
        )

        svc_entity = [
            entity
            for entity in resp.summary.entity_graph.entities
            if entity.desired_state_id == ds_id
        ][0]
        starting_state = svc_entity.starting_state.service
        starting_state_release_channels = {
            rc.release_channel: rc for rc in starting_state.release_channels
        }
        desired_state = svc_entity.desired_state.service
        for desired_release_channel_state in desired_state.release_channels:
            release_channel = desired_release_channel_state.release_channel
            starting_release_channel_state = starting_state_release_channels.get(
                release_channel
            )
            print(f"release channel: {release_channel}")
            if starting_release_channel_state:  # can be None on first deployment
                print("starting state:")
                for version in starting_release_channel_state.versions:
                    config_resp = client.service_manager.GetMaterializedConfig(
                        GetMaterializedConfigReq(
                            application=desired_state.application,
                            service=desired_state.service,
                            version=version.version,
                        )
                    )
                    svc_instance_config = [
                        cfg
                        for cfg in config_resp.compiled_service_instance_configs
                        if cfg.release_channel == release_channel
                    ][0]
                    for param in svc_instance_config.parameter_values:
                        print_param_value(param)
            print("desired state:")
            assert (
                len(desired_release_channel_state.versions) == 1
            ), "can only have one desired version"
            desired_version = desired_release_channel_state.versions[0]
            config_resp = client.service_manager.GetMaterializedConfig(
                GetMaterializedConfigReq(
                    application=desired_state.application,
                    service=desired_state.service,
                    version=desired_version.version,
                )
            )
            svc_instance_config = [
                cfg
                for cfg in config_resp.compiled_service_instance_configs
                if cfg.release_channel == release_channel
            ][0]
            for param in svc_instance_config.parameter_values:
                print_param_value(param)


if __name__ == "__main__":
    main()
