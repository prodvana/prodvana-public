import argparse
import subprocess
import sys
from typing import Optional

from prodvana.linters.check_commit_email_common import check_email


def git_config_email() -> Optional[str]:
    try:
        return (
            subprocess.check_output(["git", "config", "user.email"])
            .strip()
            .decode("utf-8")
        )
    except subprocess.CalledProcessError as e:
        if e.returncode == 1:
            return None  # not set
        raise


def main() -> None:
    ap = argparse.ArgumentParser("check-commit-email-pre-commit")
    ap.parse_args()
    from_git_config = git_config_email()
    assert (
        from_git_config
    ), "git config user.email not set. See https://git-scm.com/book/en/v2/Getting-Started-First-Time-Git-Setup."
    check_email(from_git_config)


if __name__ == "__main__":
    try:
        main()
    except AssertionError as e:
        print(
            e
        )  # pretty print error with no stack trace to stdout, which is what trunk wants
        sys.exit(1)
