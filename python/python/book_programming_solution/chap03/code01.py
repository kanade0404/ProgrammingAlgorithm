def main():
    v, n = map(int, input().split())
    nums = [int(input()) for _ in range(n)]
    is_exist = False
    for i in nums:
        if i == v:
            is_exist = True
            break
    print(is_exist)


if __name__ == '__main__':
    main()
