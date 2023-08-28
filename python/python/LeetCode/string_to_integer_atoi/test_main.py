import pytest
from .main import Solution
test_cases = [
    ("42", 42), ("   -42", -42), ("-91283472332", -2147483648),
    ("", 0), (" ", 0), ("+1", 1), ("words and 987", 0), ("00000-42a1234", 0), ("   +0 123", 0), ("  +  413", 0)
]
@pytest.mark.parametrize('test_case,expected', test_cases)
def test_my_atoi(test_case, expected):
    assert Solution().myAtoi(test_case) == expected
