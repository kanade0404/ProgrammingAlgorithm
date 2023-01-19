def main():
    n, v = map(int, input().split())
    a = [int(input()) for _ in range(n)]
    result = -1
    for i in a:
        if i == v and i > result:
            result = i
    print(result)


if __name__ == '__main__':
    main()
