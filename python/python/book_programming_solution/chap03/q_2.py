def main():
    n, v = map(int, input().split())
    a = [int(input()) for _ in range(n)]
    result = 0
    for i in a:
        if i == v:
            result += 1
    print(result)


if __name__ == '__main__':
    main()
