import pytest
from .main import Solution

test_cases = [
    ([1, 2, 3, 4], [1, 3, 6, 10]),
    ([1, 1, 1, 1, 1], [1, 2, 3, 4, 5]),
    ([3, 1, 2, 10,1], [3, 4, 6, 16, 17])
]
@pytest.mark.parametrize("input,expected", test_cases)
def test_main(input, expected):
     assert Solution().runningSum(input) == expected
