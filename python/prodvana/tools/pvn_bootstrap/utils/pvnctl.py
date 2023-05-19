import json
import subprocess
import tempfile
from typing import Any, Dict, Mapping, Optional, Sequence

import yaml

from prodvana.tools.pvn_bootstrap.utils import pulumi

SUPPORT_ORG_ID = "org-f12f30336b914efb9603a460f8412dea"


class Pvnctl:
    PVNCTL = "pvnctl"

    @classmethod
    def add_context(cls, context: str, addr: str) -> None:
        subprocess.check_call(
            [
                cls.PVNCTL,
                "auth",
                "context",
                "add",
                context,
                addr,
            ]
        )

    def __init__(self, context: str) -> None:
        self.context = context

    def pvnctl_cmd(self, args: Sequence[str]) -> Sequence[str]:
        cmd = [self.PVNCTL, "--context", self.context]
        cmd.extend(args)
        return cmd

    def is_authed(self) -> bool:
        return (
            subprocess.call(
                self.pvnctl_cmd(["runtimes", "list"]),
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL,
            )
            == 0
        )

    def auth_with_token(self, token: str) -> None:
        args = ["auth", "token", token]
        subprocess.check_call(self.pvnctl_cmd(args))

    def auth_login(self) -> None:
        subprocess.check_call(self.pvnctl_cmd(["auth", "login"]))

    def auth_login_if_needed(self) -> None:
        if not self.is_authed():
            self.auth_login()

    def add_runtime(
        self, name: str, endpoint: str, cert: str, token: str, service_account: str
    ) -> str:
        return subprocess.check_output(
            self.pvnctl_cmd(
                [
                    "runtimes",
                    "add-k8s",
                    name,
                    "--endpoint",
                    endpoint,
                    "--ca-cert",
                    cert,
                    "--token",
                    token,
                    "--service-account",
                    service_account,
                ]
            ),
            encoding="utf-8",
        ).strip()

    def configure_runtime_extensions(
        self,
        name: str,
        enable_datadog: bool = True,
        disable_istio: bool = True,
        disable_flagger: bool = True,
    ) -> str:
        config = {
            "datadog": {
                "enabled": enable_datadog,
            },
            "disableIstio": disable_istio,
            "disableFlagger": disable_flagger,
        }
        with tempfile.NamedTemporaryFile() as f:
            f.write(yaml.safe_dump(config).encode("utf-8"))
            f.flush()
            return subprocess.check_output(
                self.pvnctl_cmd(
                    [
                        "runtimes",
                        "extensions",
                        "configure",
                        name,
                        "--input-file-format=yaml",
                        "--no-prompt",
                        f"--input={f.name}",
                    ]
                ),
                encoding="utf-8",
            ).strip()

    def add_runtime_from_management_cluster_output(
        self, name: str, output: pulumi.ManagementClusterOutput
    ) -> str:
        return self.add_runtime(
            name=name,
            endpoint=output["cluster_k8s_endpoint"],
            cert=output["prodvana_service_account_ca_crt"],
            token=output["prodvana_service_account_token"],
            service_account=output["prodvana_service_account_name"],
        )

    def add_runtime_from_user_cluster_output(
        self, name: str, output: pulumi.UserClusterOutput
    ) -> str:
        return self.add_runtime(
            name=name,
            endpoint=output["k8s_endpoint"],
            cert=output["prodvana_service_account_ca_crt"],
            token=output["prodvana_service_account_token"],
            service_account=output["prodvana_service_account_name"],
        )

    def configure_release_channel(
        self,
        app: str,
        name: str,
        runtime: str,
        default_env: Mapping[str, str],
        capabilities_refs: Mapping[str, str],
        preconditions: Sequence[object] | None,
        namespace: str | None,
    ) -> None:
        config = yaml.safe_load(
            subprocess.check_output(
                self.pvnctl_cmd(
                    ["applications", "get-config", app, "--original-config"]
                )
            )
        )
        rc_configs = [c for c in config["releaseChannels"] if c["name"] == name]
        if rc_configs:
            assert len(rc_configs) == 1
            rc_config = rc_configs[0]
        else:
            rc_config = {"name": name}
            config["releaseChannels"].append(rc_config)
        rc_config["policy"] = {
            "defaultEnv": {k: {"value": v} for k, v in default_env.items()},
        }
        if not any(r["runtime"] == runtime for r in rc_config.get("runtimes", [])):
            rc_config.setdefault("runtimes", [])
            runtimeCfg: Dict[str, Any] = {"runtime": runtime}
            if namespace is not None:
                runtimeCfg["containerOrchestration"] = {
                    "k8s": {
                        "namespace": namespace,
                    }
                }
            rc_config["runtimes"].append(runtimeCfg)

        for cap in config["capabilities"]:
            new_ref = capabilities_refs[cap["name"]]
            cap.setdefault("perReleaseChannel", [])
            rc_caps = [
                c for c in cap["perReleaseChannel"] if c["releaseChannel"] == name
            ]
            if rc_caps:
                assert len(rc_caps) == 1
                rc_caps[0] = new_ref
            else:
                cap["perReleaseChannel"].append(
                    {"releaseChannel": name, "ref": {"name": new_ref}}
                )
        if preconditions:
            rc_config["preconditions"] = preconditions

        with tempfile.NamedTemporaryFile() as f:
            f.write(yaml.safe_dump(config).encode("utf-8"))
            f.flush()
            subprocess.check_call(
                self.pvnctl_cmd(
                    [
                        "applications",
                        "edit",
                        app,
                        "--input-file-format=yaml",
                        "--no-prompt",
                        f"--input={f.name}",
                    ]
                )
            )

    def create_service(
        self,
        config_path: str,
    ) -> None:
        subprocess.check_call(
            self.pvnctl_cmd(
                [
                    "services",
                    "--app=default",
                    "create",
                    "--input",
                    config_path,
                ]
            )
        )

    def push_existing(
        self,
        service: str,
    ) -> None:
        subprocess.check_call(
            self.pvnctl_cmd(
                [
                    "--editor=/bin/true",  # hack to not do anything
                    "services",
                    "--app=default",
                    "push",
                    "--no-prompt",
                    service,
                ]
            )
        )

    # useful for configuring existing service for new release channel
    def configure_existing(
        self,
        service: str,
    ) -> str:
        return subprocess.check_output(
            self.pvnctl_cmd(
                [
                    "--editor=/bin/true",  # hack to not do anything
                    "services",
                    "--app=default",
                    "push",
                    "--no-push",
                    "--no-prompt",
                    service,
                ]
            ),
            encoding="utf-8",
        ).strip()

    def configure_service(
        self,
        config_path: str,
    ) -> str:
        return subprocess.check_output(
            self.pvnctl_cmd(
                [
                    "services",
                    "--app=default",
                    "primitives",
                    "configure-proto-experimental",
                    config_path,
                ]
            ),
            encoding="utf-8",
        ).strip()

    def push_service_instance(
        self,
        service: str,
        release_channel: str,
        version: str = "latest",
    ) -> None:
        subprocess.check_call(
            self.pvnctl_cmd(
                [
                    "services",
                    "--app=default",
                    "primitives",
                    "push",
                    service,
                    release_channel,
                    f"--version={version}",
                ]
            )
        )

    def health_check(self) -> bool:
        return (
            subprocess.call(
                self.pvnctl_cmd(["healthcheck"]),
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL,
            )
            == 0
        )

    def list_services(self) -> Mapping[str, Any]:
        output = subprocess.check_output(
            self.pvnctl_cmd(["-f", "json", "services", "--app=default", "list"])
        ).strip()
        services = json.loads(output)
        svc_by_name: Dict[str, Dict[str, Any]] = {}
        for service in services:
            svc_by_name.setdefault(service["service"], {})
            svc_by_name[service["service"]][service["releaseChannel"]] = service
        return svc_by_name


_BOOTSTRAP_PVNCTL: Optional[Pvnctl] = None


def bootstrap_pvnctl() -> Pvnctl:
    global _BOOTSTRAP_PVNCTL
    if _BOOTSTRAP_PVNCTL is None:
        Pvnctl.add_context("bootstrap", "35.232.161.117:5000")
        _BOOTSTRAP_PVNCTL = Pvnctl("bootstrap")
    return _BOOTSTRAP_PVNCTL


_PRIME_PVNCTL: Optional[Pvnctl] = None


def prime_pvnctl() -> Pvnctl:
    global _PRIME_PVNCTL
    if _PRIME_PVNCTL is None:
        Pvnctl.add_context("prime", "api.prime.prodvana.io:443")
        _PRIME_PVNCTL = Pvnctl("prime")
    return _PRIME_PVNCTL
