import sys


def main():
    n, k = map(int, input().split())
    a = [int(input()) for _ in range(n)]
    b = [int(input()) for _ in range(n)]
    min = sys.maxsize
    for i in range(n):
        for j in range(n):
            if a[i] + b[j] < k:
                continue
            if a[i] + b[j] < min:
                min = a[i] + b[j]
    print(min)


if __name__ == '__main__':
    main()
