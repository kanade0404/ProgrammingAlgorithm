#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'timeConversion' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#

def convert(h: int, am_pm: str) -> int:
    if am_pm == "AM" and h >= 12:
        return 0 if h == 12 else h % 12
    if am_pm == "PM" and h <= 12:
        return h if h == 12 else 12 + h
    return h

def timeConversion(s):
    am_pm = s[-2:]
    hour, minute, second = s[:-2].split(":")
    h = convert(int(hour), am_pm)
    print(h,  hour, minute, second, am_pm)
    return "{}:{}:{}".format(str(h) if h >= 10 else "0{}".format(h), minute, second)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)

    fptr.write(result + '\n')

    fptr.close()
