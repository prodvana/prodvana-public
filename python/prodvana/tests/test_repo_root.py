import os

from prodvana.repo_root import REPO_ROOT


def test_repo_root() -> None:
    assert os.path.exists(os.path.join(REPO_ROOT))
    assert os.path.exists(os.path.join(REPO_ROOT, "python"))
    assert os.path.exists(os.path.join(REPO_ROOT, "python", "prodvana"))
