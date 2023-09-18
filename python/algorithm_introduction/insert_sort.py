from typing import List

def insert_sort(arr: List[int]) -> List[int]:
    num = len(arr)
    for j in range(2, num):
        key = arr[j]
        i = j - 1
        while i > 0 and arr[i] > key:
            arr[i + 1] = arr[i]
            i = i - 1
        arr[i + 1] = key
    return arr
