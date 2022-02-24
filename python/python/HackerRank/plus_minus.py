#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'plusMinus' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#

def plusMinus(arr):
    total_len = len(arr)
    print('{:.6f}'.format(round(len(list(filter(lambda x: x > 0, arr))) / total_len, 6)))
    print('{:.6f}'.format(round(len(list(filter(lambda x: x < 0, arr))) / total_len, 6)))
    print('{:.6f}'.format(round(len(list(filter(lambda x: x == 0, arr))) / total_len, 6)))

if __name__ == '__main__':
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
