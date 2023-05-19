import argparse

from prodvana.tools.pvn_bootstrap.utils import pulumi, pvnctl


def main() -> None:
    ap = argparse.ArgumentParser("init_bootstrap_cluster")
    ap.parse_args()
    output = pulumi.get_management_cluster_output(
        stack=pulumi.ManagementClusterStacks.PRODVANA_PRIME
    )
    pvnctl.bootstrap_pvnctl().add_runtime_from_management_cluster_output(
        "prodvana-prime", output
    )


if __name__ == "__main__":
    main()
