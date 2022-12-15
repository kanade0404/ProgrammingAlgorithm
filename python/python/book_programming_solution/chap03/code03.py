import sys


def main():
    n = int(input())
    nums = [int(input()) for _ in range(n)]
    min_val = sys.maxsize
    for i in nums:
        if i < min_val:
            min_val = i
    print(min_val)


if __name__ == '__main__':
    main()
