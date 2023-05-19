import contextlib
import json
import os
import subprocess
from typing import Any, Dict, Iterator, Literal, Optional, TypedDict, cast

import yaml

from prodvana.repo_root import REPO_ROOT
from prodvana.tools.pvn_bootstrap.utils.constants import LEGACY_NAMESPACE_NAME_ORG_IDS

PULUMI = os.path.join(REPO_ROOT, "tools", "pvn-pulumi")
MANAGEMENT_CLUSTER_DIR = os.path.join(REPO_ROOT, "pulumi/projects/management-cluster")
USER_CLUSTER_DIR = os.path.join(REPO_ROOT, "pulumi/projects/user-cluster")
SUPPORT_CLUSTER_DIR = os.path.join(REPO_ROOT, "pulumi/projects/support-cluster")
SECRETS_PROVIDER = "gcpkms://projects/pulumi-admin-58265/locations/global/keyRings/pulumi-keyring-15a1d77/cryptoKeys/pulumi-secret-provider-key-99ab97a"

ManagementClusterOutput = TypedDict(
    "ManagementClusterOutput",
    {
        "cluster_k8s_endpoint": str,
        "gadmin-lb-ip-address": Optional[str],
        "gadmin-lb-ip-name": Optional[str],
        "gapi-lb-ip-address": str,
        "gapi-lb-ip-name": str,
        "prodvana_config_id": str,
        "prodvana_config_path": str,
        "prodvana_service_account_ca_crt": str,
        "prodvana_service_account_name": str,
        "prodvana_service_account_token": str,
        "pulumi_secret_key_id": str,
        "pulumi_service_account_email": str,
        "pulumi_state_bucket_url": str,
        "ssh_public_key": str,
        "org_id": str,
        "org_name": str,
        "org_slug": str,
        "namespace": str,
        "env": Literal["staging", "prod"],
    },
)

UserClusterOutput = TypedDict(
    "UserClusterOutput",
    {
        "k8s_endpoint": str,
        "prodvana_service_account_ca_crt": str,
        "prodvana_service_account_name": str,
        "prodvana_service_account_token": str,
    },
)

SupportClusterOutput = TypedDict(
    "SupportClusterOutput",
    {
        "k8s_endpoint": str,
        "prodvana_service_account_ca_crt": str,
        "prodvana_service_account_name": str,
        "prodvana_service_account_token": str,
        "prodvana_spanner_resource": str,
    },
)


class ManagementClusterStacks:
    PRODVANA_PRIME = "management_cluster_prodvana_prime"


@contextlib.contextmanager
def login(login_url: str, pulumi_dir: str) -> Iterator[None]:
    subprocess.check_call([PULUMI, "login", login_url], cwd=pulumi_dir)
    try:
        yield
    finally:
        subprocess.check_call([PULUMI, "logout"], cwd=pulumi_dir)


def get_output(pulumi_dir: str, stack: str) -> Dict[str, Any]:
    output = subprocess.check_output(
        [PULUMI, "--stack", stack, "stack", "output", "--json", "--show-secrets"],
        cwd=pulumi_dir,
    )
    return cast(Dict[str, Any], json.loads(output))


def get_management_cluster_output(stack: str) -> ManagementClusterOutput:
    return cast(
        ManagementClusterOutput,
        get_output(MANAGEMENT_CLUSTER_DIR, stack=stack),
    )


def get_user_cluster_output(stack: str) -> UserClusterOutput:
    return cast(
        UserClusterOutput,
        get_output(USER_CLUSTER_DIR, stack=stack),
    )


def get_support_cluster_output(stack: str) -> SupportClusterOutput:
    return cast(
        SupportClusterOutput,
        get_output(SUPPORT_CLUSTER_DIR, stack=stack),
    )


def create_management_cluster_stack(
    org_slug: str, org_id: str, org_name: str, env: str, sso_conn: str | None
) -> str:
    """
    Create management cluster stack config, return stack name.
    """
    stack_name = f"management_cluster_{org_slug}"
    if org_id in LEGACY_NAMESPACE_NAME_ORG_IDS:
        namespace = (
            f"pvn-{org_id}" if org_id.startswith("org-") else f"pvn-org-{org_id}"
        )
    else:
        namespace = f"pvn-default-rc-{org_id}"

    config: Dict[str, Any] = {
        "secretsprovider": SECRETS_PROVIDER,
        "encryptedkey": "CiUACFMnKn30hyB1mmzp6gZF6qG+P+rm1gMybDJYofRgCXeQ5hcGEkkACLzZ8KxammLUdF5tqfKjC51daHs6kOLiYxvdEfxbS841xaxNVhKPMXTHRqVC2v27NlIYodZtB0TntpVuqURvqdLduwkk8DGP",
        "config": {
            "gcp:project": "pvn-mgmt",
            "management-cluster:active": True,
            "management-cluster:env": env,
            "management-cluster:instance_type": "e2-standard-4",
            "management-cluster:node_count_per_zone": "2",
            "management-cluster:nodepool_zones": [
                "us-central1-a",
                "us-central1-c",
            ],
            "management-cluster:org_id": org_id,
            "management-cluster:org_name": org_name,
            "management-cluster:org_slug": org_slug,
            "management-cluster:region": "us-central1",
            "management-cluster:namespace": namespace,
            "management-cluster:writeback_repo": org_slug,
            "management-cluster:writeback_repo_org": "prodvana-user-configs",
        },
    }
    if sso_conn:
        config["config"]["management-cluster:auth0_connections"] = [sso_conn]
    with open(
        os.path.join(MANAGEMENT_CLUSTER_DIR, f"Pulumi.{stack_name}.yaml"), "w"
    ) as f:
        f.write(yaml.safe_dump(config, sort_keys=False))
    return stack_name


def create_user_cluster_stack_gcp(normalized_org_name: str, org_id: str) -> str:
    """
    Create user cluster stack config, return stack name.
    """
    stack_name = f"user_cluster_{normalized_org_name}"
    with open(os.path.join(USER_CLUSTER_DIR, f"Pulumi.{stack_name}.yaml"), "w") as f:
        f.write(
            yaml.safe_dump(
                {
                    "secretsprovider": SECRETS_PROVIDER,
                    "encryptedkey": "CiUACFMnKn30hyB1mmzp6gZF6qG+P+rm1gMybDJYofRgCXeQ5hcGEkkACLzZ8KxammLUdF5tqfKjC51daHs6kOLiYxvdEfxbS841xaxNVhKPMXTHRqVC2v27NlIYodZtB0TntpVuqURvqdLduwkk8DGP",
                    "config": {
                        "user-cluster:gcp_project": "pvn-infra",
                        "user-cluster:cluster_provider": "gcp",
                        "user-cluster:instance_type": "e2-standard-4",
                        "user-cluster:name": f"user-{org_id}",
                        "user-cluster:node_count_per_zone": "2",
                        "user-cluster:nodepool_zones": [
                            "us-central1-a",
                            "us-central1-c",
                        ],
                        "user-cluster:region": "us-central1",
                    },
                },
                sort_keys=False,
            )
        )
    return stack_name


def create_user_cluster_stack_eks(normalized_org_name: str, org_id: str) -> str:
    """
    Create user cluster stack config, return stack name.
    """
    stack_name = f"user_cluster_{normalized_org_name}"
    with open(os.path.join(USER_CLUSTER_DIR, f"Pulumi.{stack_name}.yaml"), "w") as f:
        f.write(
            yaml.safe_dump(
                {
                    "secretsprovider": SECRETS_PROVIDER,
                    "encryptedkey": "CiUACFMnKn30hyB1mmzp6gZF6qG+P+rm1gMybDJYofRgCXeQ5hcGEkkACLzZ8KxammLUdF5tqfKjC51daHs6kOLiYxvdEfxbS841xaxNVhKPMXTHRqVC2v27NlIYodZtB0TntpVuqURvqdLduwkk8DGP",
                    "config": {
                        "user-cluster:cluster_provider": "eks",
                        "user-cluster:aws_account_id": "717408832317",
                        "user-cluster:aws_assume_role_arn": "arn:aws:iam::717408832317:role/OrganizationAccountAccessRole",
                        "user-cluster:instance_type": "m5a.xlarge",
                        "user-cluster:name": f"user-{org_id}",
                        "user-cluster:node_count_per_zone": "2",
                        "user-cluster:nodepool_zones": [
                            "us-west-2a",
                            "us-west-2b",
                            "us-west-2c",
                        ],
                        "user-cluster:region": "us-west-2",
                    },
                },
                sort_keys=False,
            )
        )
    return stack_name


create_user_cluster_stack = {
    "eks": create_user_cluster_stack_eks,
    "gcp": create_user_cluster_stack_gcp,
}
