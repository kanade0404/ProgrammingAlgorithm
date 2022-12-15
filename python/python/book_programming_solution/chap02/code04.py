import sys


def main():
    n = int(input())
    result = sys.maxsize
    for _ in range(n):
        x, y = map(int, input().split())
        if result > abs(x - y):
            result = abs(x - y)
    print(result)


if __name__ == '__main__':
    main()
