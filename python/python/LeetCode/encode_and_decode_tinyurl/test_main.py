import pytest
from .main import Codec

test_cases = [
    "https://leetcode.com/problems/design-tinyurl",
]


@pytest.mark.parametrize('test_case', test_cases)
def test_main(test_case):
    codec = Codec()
    assert codec.decode(codec.encode(test_case)) == test_case
