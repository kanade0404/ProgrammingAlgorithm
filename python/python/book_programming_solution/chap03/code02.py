def main():
    v, n = map(int, input().split())
    nums = [int(input()) for _ in range(n)]
    found_id = -1
    for i, x in enumerate(nums):
        if v == x:
            found_id = i
            break
    print(found_id)


if __name__ == '__main__':
    main()
