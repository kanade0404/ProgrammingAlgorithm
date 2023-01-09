def main():
    n, k = map(int, input().split())
    a = [int(input()) for _ in range(n)]
    is_exist = False
    for bit in range(2 ** n):
        sum = 0
        for i in range(n):
            if bit & (1 << i):
                sum += a[n - 1 - i]
        if sum == k:
            is_exist = True
            break
    print("Yes" if is_exist else "No")


if __name__ == '__main__':
    main()
