import sys


def main():
    n = int(input())
    a = [int(input()) for _ in range(n)]
    first_min, second_min = sys.maxsize, sys.maxsize
    for i in a:
        if i < second_min:
            second_min = i
        if second_min < first_min:
            tmp = first_min
            first_min = second_min
            second_min = tmp
    print(second_min)


if __name__ == '__main__':
    main()
