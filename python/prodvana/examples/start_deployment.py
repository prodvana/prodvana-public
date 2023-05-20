import argparse

from prodvana.client import Client, make_channel
from prodvana.proto.prodvana.application.application_manager_pb2 import (
    GetApplicationReq,
)
from prodvana.proto.prodvana.common_config.parameters_pb2 import ParameterValue
from prodvana.proto.prodvana.desired_state.manager_pb2 import SetDesiredStateReq
from prodvana.proto.prodvana.desired_state.model.desired_state_pb2 import (
    ServiceInstanceState,
    ServiceState,
    State,
    Version,
)
from prodvana.proto.prodvana.service.service_manager_pb2 import (
    ApplyParametersReq,
    GetServiceConfigReq,
    ServiceConfigVersionReference,
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
    ap.add_argument("app", help="Application that contains the service")
    ap.add_argument("service", help="Service to deploy")
    ap.add_argument("param_name", help="Parameter name to update")
    ap.add_argument("param_value", help="Parameter value to set")
    args = ap.parse_args()
    app = args.app
    service = args.service
    param_name = args.param_name
    param_value = args.param_value
    with make_channel(org=args.org, api_token=args.api_token) as channel:
        client = Client(channel=channel)

        # take the latest config
        config_resp = client.service_manager.GetServiceConfig(
            GetServiceConfigReq(application=app, service=service)
        )

        # validate that the requested parameter exists and is a docker image parameter (the only one supported by this example)
        param_defs = {param.name: param for param in config_resp.config.parameters}
        assert param_name in param_defs
        assert param_defs[param_name].docker_image

        # create a new service version using the service config and requested parameter
        apply_resp = client.service_manager.ApplyParameters(
            ApplyParametersReq(
                service_config_version=ServiceConfigVersionReference(
                    application=app,
                    service=service,
                    service_config_version=config_resp.config_version,
                ),
                parameters=[
                    ParameterValue(name=param_name, docker_image_tag=param_value),
                ],
            )
        )

        # get list of release channels so we can construct desired state
        app_resp = client.application_manager.GetApplication(
            GetApplicationReq(application=app)
        )

        # construct desired state and set it, which causes convergence to begin
        desired_state = State(
            service=ServiceState(
                service=service,
                application=app,
                release_channels=[
                    ServiceInstanceState(
                        release_channel=rc.name,
                        versions=[
                            Version(
                                version=apply_resp.version,
                            ),
                        ],
                    )
                    for rc in app_resp.application.config.release_channels
                ],
            )
        )
        ds_resp = client.desired_state_manager.SetDesiredState(
            SetDesiredStateReq(desired_state=desired_state)
        )

        # desired state id is unique for each call to SetDesiredState and can be used to get convergence status, submit manual approval, etc.
        print(ds_resp.desired_state_id)


if __name__ == "__main__":
    main()
