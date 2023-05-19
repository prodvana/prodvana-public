import argparse
import logging
import time
from typing import Callable

from prodvana.tools.pvn_bootstrap.utils import constants, pulumi, pvnadminctl, pvnctl

PVN_STAGING_RELEASE_CHANNEL = "org-c00f892272374b958e933c2e119b8dd7"
PVN_SCAN_RELEASE_CHANNEL = "org-17aac462ee7e4488b1d236ca8c4a839b"

RELEASE_CHANNEL_ENV_PRECONDITIONS = {
    "staging": None,
    "prod": [
        {"manualApproval": {}},
        {
            "releaseChannelStable": {"releaseChannel": PVN_STAGING_RELEASE_CHANNEL},
        },
        {
            "releaseChannelStable": {"releaseChannel": PVN_SCAN_RELEASE_CHANNEL},
        },
    ],
}


def wait_for(
    condition: Callable[[], bool], condition_message: str, timeout_seconds: float = 180
) -> None:
    start = time.time()
    first = True
    while time.time() - start <= timeout_seconds:
        if not first:
            logging.info("Waiting for condition '%s' to be true", condition_message)
            time.sleep(5)
        first = False
        if condition():
            return
    raise Exception("Timed out waiting for condition")


def main() -> None:
    logging.basicConfig(level=logging.INFO)
    ap = argparse.ArgumentParser("init_user_cluster")
    ap.add_argument(
        "--management-stack", required=True, help="Management user cluster stack"
    )
    ap.add_argument(
        "--skip-org-init",
        action="store_true",
        help="Skip pvnadminctl commands to initialize organization.",
    )
    ap.add_argument(
        "--skip-services-init",
        action="store_true",
        help="Skip pushing services to the organization management cluster.",
    )
    ap.add_argument("--user-stack", help="User cluster stack, omit if not created")
    args = ap.parse_args()
    management_output = pulumi.get_management_cluster_output(
        stack=args.management_stack,
    )
    org_id = management_output["org_id"]
    org_name = management_output["org_name"]
    org_slug = management_output["org_slug"]
    env = management_output["env"]
    namespace = management_output["namespace"]
    assert env in ("staging", "prod")

    logging.info("Authenticating if needed")
    pvnctl.prime_pvnctl().auth_login_if_needed()

    if not args.skip_org_init:
        logging.info("Creating organization")
        generated_org_id = pvnadminctl.pvnadminctl_for_env(env).create_organization(
            org_name, slug=org_slug, org_id=org_id
        )
        assert org_id == generated_org_id

    logging.info("Adding user management cluster to Prodvana Prime")
    pvnctl.prime_pvnctl().add_runtime_from_management_cluster_output(
        org_id, management_output
    )
    pvnctl.prime_pvnctl().configure_runtime_extensions(
        org_id,
        enable_datadog=True,
        disable_istio=True,
        disable_flagger=True,
    )
    if org_id.startswith("org-"):
        rc = org_id
    else:
        rc = f"org-{org_id}"

    base_domain = {
        "prod": "prodvana.io",
        "staging": "staging.prodvana.io",
    }[env]
    ext_addr = f"api.{org_slug}.{base_domain}:443"
    logging.info("User's apiserver addr is %s", ext_addr)

    # set env again with apiserver ip
    pvnctl.prime_pvnctl().configure_release_channel(
        name=rc,
        app="default",
        runtime=org_id,
        default_env={
            "PVN_CONFIG": management_output["prodvana_config_path"],
            "PVN_ENVIRONMENT": env,
            "APISERVER_INTERNAL_ADDR": f"apiserver.pvn-{rc}:5000"
            if org_id in constants.LEGACY_NAMESPACE_NAME_ORG_IDS
            else f"apiserver.pvn-default-rc-{rc}:5000",
            "APISERVER_EXTERNAL_ADDR": ext_addr,
            "AGENT_INTERACTION_EXTERNAL_ADDRESS": f"agt.{org_slug}.{base_domain}:443",
            "UI_EXTERNAL_HTTP_ADDR": f"https://{org_slug}.{base_domain}",
        },
        capabilities_refs={
            "auth0-team-management": f"auth0-team-management-{env}",
            "spanner": f"spanner-{env}",
            "honeycomb": f"honeycomb-{env}",
        },
        preconditions=RELEASE_CHANNEL_ENV_PRECONDITIONS[env],
        namespace=namespace,
    )
    if not args.skip_services_init:
        for service in (
            "agent-interaction",
            "apiserver",
            "cache-warmer",
            "dynamic-delivery",
            "insights-generator",
            "ui",
        ):
            pvnctl.prime_pvnctl().push_existing(service)

    pvnctl.Pvnctl.add_context(org_slug, ext_addr)
    user_pvnctl = pvnctl.Pvnctl(org_slug)
    wait_for(user_pvnctl.health_check, "user's apiserver healthy")

    if args.user_stack:
        user_pvnctl.auth_login_if_needed()
        user_output = pulumi.get_user_cluster_output(
            stack=args.user_stack,
        )
        logging.info(
            "Adding user cluster IP %s to user management cluster",
            user_output["k8s_endpoint"],
        )
        user_pvnctl.add_runtime_from_user_cluster_output(
            name="default", output=user_output
        )


if __name__ == "__main__":
    main()
