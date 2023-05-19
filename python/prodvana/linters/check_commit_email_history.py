import argparse
import subprocess
import sys
from typing import Iterator

from prodvana.linters.check_commit_email_common import check_email


def get_commit_email(ref: str) -> str:
    return (
        subprocess.check_output(["git", "log", r"--format=%ae", f"{ref}^!"])
        .strip()
        .decode("utf-8")
    )


def get_parent_commits(ref: str) -> Iterator[str]:
    i = 1
    while True:
        try:
            yield subprocess.check_output(
                ["git", "rev-parse", f"{ref}^{i}"]
            ).strip().decode("utf-8")
        except subprocess.CalledProcessError:
            return
        i += 1


def validate_commit_author(ref: str) -> None:
    email = get_commit_email(ref=ref)
    if "mergify" in email:
        # if this is a mergify commit, look at the parent commit(s) instead as this is likely a merge commit
        for commit in get_parent_commits(ref=ref):
            validate_commit_author(ref=commit)
        return
    check_email(email)


def main() -> int:
    ap = argparse.ArgumentParser("check-commit-email-history")
    ap.parse_args()
    validate_commit_author(ref="HEAD")
    return 0


if __name__ == "__main__":
    sys.exit(main())
