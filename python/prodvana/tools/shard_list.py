import argparse
import hashlib
import sys
from typing import Sequence


def shard(items: Sequence[str], shard_idx: int, total: int) -> Sequence[str]:
    test_hashes = {t: hashlib.md5(t.encode("utf-8")).digest() for t in items}
    items = sorted(items, key=lambda t: test_hashes[t])
    sharded = []
    for idx, item in enumerate(items):
        if idx % total == shard_idx:
            sharded.append(item)
    return sharded


def main() -> None:
    ap = argparse.ArgumentParser("shard_list")
    ap.add_argument("--index", required=True, type=int)
    ap.add_argument("--total", required=True, type=int)
    args = ap.parse_args()
    assert 0 <= args.index < args.total
    tests = [t.strip() for t in sys.stdin.readlines() if t.strip()]
    sharded_tests = shard(tests, shard_idx=args.index, total=args.total)
    for test in sharded_tests:
        print(test)


if __name__ == "__main__":
    main()
