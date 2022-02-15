from typing import List
SYMBOL = {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000
}


def romanToInt(s: str) -> int:
    str_list: List[int] = list(s)
    result = 0
    for i, x in enumerate(str_list):
        result = result + SYMBOL[x] - SYMBOL[str_list[i - 1]] * 2 \
            if result != 0 and SYMBOL[str_list[i - 1]] < SYMBOL[x] else result + SYMBOL[x]
    return result
