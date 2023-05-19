import subprocess
from typing import Dict, Literal, Optional, Sequence


class Pvnadminctl:
    PVNADMINCTL = "pvnadminctl"

    @classmethod
    def add_context(cls, context: str, addr: str) -> None:
        subprocess.check_call(
            [
                cls.PVNADMINCTL,
                "auth",
                "context",
                "add",
                context,
                addr,
            ]
        )

    def __init__(self, context: str) -> None:
        self.context = context

    def pvnadminctl_cmd(self, args: Sequence[str]) -> Sequence[str]:
        cmd = [self.PVNADMINCTL, "--context", self.context]
        cmd.extend(args)
        return cmd

    def create_organization(
        self, name: str, slug: str, org_id: Optional[str] = None
    ) -> str:
        """
        Create a new organization, return its ID.
        """
        cmd = ["org", "create", "--slug", slug, name]
        if org_id:
            cmd.extend(["--id", org_id])
        id = subprocess.check_output(self.pvnadminctl_cmd(cmd), encoding="utf-8")
        return id.strip()


_PVNADMINCTLS: Dict[str, Pvnadminctl] = {}


def pvnadminctl_for_env(env: Literal["staging", "prod"]) -> Pvnadminctl:
    if env not in _PVNADMINCTLS:
        admin_ip = f"admin-server.pvn-support-rc-{env}.svc"
        Pvnadminctl.add_context(env, f"{admin_ip}:6000")
        _PVNADMINCTLS[env] = Pvnadminctl(env)
    return _PVNADMINCTLS[env]
