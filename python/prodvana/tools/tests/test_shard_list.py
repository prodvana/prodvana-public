import pytest

from prodvana.tools.shard_list import shard


@pytest.mark.parametrize("size", range(20))
def test_shard(size: int) -> None:
    items = [str(i) for i in range(size)]
    items_set = set(items)
    total = 3
    for index in range(total):
        sharded = shard(items, shard_idx=index, total=total)
        if size >= total:
            assert sharded
        for item in sharded:
            assert item in items_set
            items_set.remove(item)
    assert len(items_set) == 0, "not all items chosen"
