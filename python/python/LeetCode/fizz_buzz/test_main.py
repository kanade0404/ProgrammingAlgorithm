import pytest
from .main import Solution

test_cases = [
    (3, ["1", "2", "Fizz"]),
    (5, ["1", "2", "Fizz", "4", "Buzz"]),
    (15, ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"])
]
@pytest.mark.parametrize("input,expected", test_cases)
def test_main(input, expected):
    assert Solution().fizzBuzz(input) == expected

