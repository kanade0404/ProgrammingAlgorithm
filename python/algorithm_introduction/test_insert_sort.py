from pytest import mark
from algorithm_introduction import insert_sort

test_cases = [
    ([5, 2, 4, 6 ,1 ,3], [1, 2, 3, 4, 5, 6]),
    ([1, 2, 3, 4, 5, 6], [1, 2, 3, 4, 5, 6]),
    ([2, 0, 1, 3, 6, 5], [1, 2, 3, 4, 5, 6]),
    ([0], [0]),
]
@mark.parametrize('test_case,expected', [])
def test_insert_sort(test_case, expected):
    assert insert_sort(test_case) == expected
