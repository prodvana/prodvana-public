import argparse
import logging
import uuid

from prodvana.tools.pvn_bootstrap.utils import pulumi, slugify


def main() -> None:
    logging.basicConfig(level=logging.INFO)
    ap = argparse.ArgumentParser("init_user_organization_pulumi_files")
    ap.add_argument("name", help="Human-readable organization name")
    ap.add_argument(
        "--org-slug",
        help="Machine-friendly name to use for things like DNS. If not provided, this will be derived from organization name.",
    )
    ap.add_argument("--org-id", help="Org ID to use, will generate one if not provided")
    ap.add_argument(
        "--target",
        default="gcp",
        help="Target user cluster type - 'gcp' (default) or 'eks'",
    )
    ap.add_argument(
        "--no-user",
        action="store_false",
        dest="user",
        help="Skip user cluster creation, only create management cluster",
    )
    ap.add_argument(
        "--env",
        required=True,
        choices=("staging", "prod"),
        help="Prodvana management cluster's environment",
    )
    ap.add_argument(
        "--sso-conn",
        help="Optional. The Auth0 connection ID of the orgnanizations SSO provider.",
    )
    ap.set_defaults(user=True)
    args = ap.parse_args()
    name: str = args.name
    slug = args.org_slug or slugify(name)
    if args.org_id:
        org_id = args.org_id
    else:
        # keep in sync with go/id/id.go
        org_id = f"org-{uuid.uuid4().hex}"

    logging.info("Writing out stack files")
    management_stack_name = pulumi.create_management_cluster_stack(
        org_slug=slug,
        org_id=org_id,
        env=args.env,
        org_name=args.name,
        sso_conn=args.sso_conn,
    )
    user_msg = ""
    user_args = ""
    if args.user:
        user_stack_name = pulumi.create_user_cluster_stack[args.target](
            normalized_org_name=slug, org_id=org_id
        )
        user_args = f"--user-stack {user_stack_name}"
        user_msg = f"""
(
    cd {pulumi.USER_CLUSTER_DIR}
    pvn-pulumi stack init {user_stack_name} --secrets-provider={pulumi.SECRETS_PROVIDER}
    pvn-pulumi --stack {user_stack_name} up
)
"""
    logging.info(
        f"""Success. Please *check in all changes* and run the following pulumi commands:
(
    cd {pulumi.MANAGEMENT_CLUSTER_DIR}
    pvn-pulumi stack init {management_stack_name} --secrets-provider={pulumi.SECRETS_PROVIDER}
    pvn-pulumi --stack {management_stack_name} up
)
{user_msg}
Then run: pvn-poetry-run python -m prodvana.tools.pvn_bootstrap.init_user_organization_pvnctl --management-stack {management_stack_name} {user_args}
"""
    )


if __name__ == "__main__":
    main()
