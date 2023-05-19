"""
Example usage: pvn-poetry-run python3 -m prodvana.tools.build_go_release_binaries prodvana/cmd/pvnctl --out=bin/release
"""
import argparse
import logging
import os
import subprocess

from prodvana.repo_root import REPO_ROOT

GO = "/usr/local/go/bin/go"


def build_go_binary(
    package: str, target_os: str, target_arch: str, out_dir: str
) -> None:
    logging.info("Building %s for %s/%s", package, target_os, target_arch)
    go_dir = os.path.join(REPO_ROOT, "go")
    assert out_dir.startswith("/"), "out_dir must be absolute"
    binary_base_name = os.path.basename(package)
    arch_dir = f"{binary_base_name}_{target_os}_{target_arch}"
    output_dir = os.path.join(out_dir, arch_dir)
    tar_name = f"{arch_dir}.tar.gz"
    destination_tar = os.path.join(out_dir, tar_name)
    if not os.path.exists(output_dir):
        os.mkdir(output_dir)

    go_build_cmd = [
        GO,
        "build",
        "-ldflags=-extldflags=-static",
        "-tags",
        "osusergo,netgo",
        "-gcflags",
        f"all=-trimpath={go_dir}",
        "-o",
        os.path.join(output_dir, binary_base_name),
        package,
    ]
    env = os.environ.copy()
    env["GOOS"] = target_os
    env["GOARCH"] = target_arch
    subprocess.check_call(go_build_cmd, cwd=go_dir, env=env)
    tar_cmd = [
        "tar",
        "-czvf",
        destination_tar,
        "-C",
        out_dir,
        arch_dir,
    ]
    print(tar_cmd)
    subprocess.check_call(tar_cmd, env=env)


def build_go_package(package: str, out_dir: str) -> None:
    for target_os, target_arch in (
        ("linux", "amd64"),
        ("linux", "arm64"),
        ("darwin", "amd64"),
        ("darwin", "arm64"),
    ):
        build_go_binary(
            package=package,
            out_dir=out_dir,
            target_os=target_os,
            target_arch=target_arch,
        )


def main() -> None:
    logging.basicConfig(level=logging.INFO)
    ap = argparse.ArgumentParser("build_go_release_binaries")
    ap.add_argument("--out", required=True, help="Output directory")
    ap.add_argument("go_packages", nargs="+", help="Go packages to build")
    args = ap.parse_args()
    os.makedirs(args.out, exist_ok=True)
    out = os.path.realpath(args.out)
    for package in args.go_packages:
        build_go_package(package, out_dir=out)


if __name__ == "__main__":
    main()
