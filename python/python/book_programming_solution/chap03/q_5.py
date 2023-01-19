def main():
    n = int(input())
    a = [int(input()) for _ in range(n)]
    all_even = True
    count = 0
    while all_even:
        for i in range(n):
            if a[i] % 2 == 0:
                a[i] = a[i]//2
            else:
                a[i] = a[i]//2 + a[i]%2
                all_even = False
        if not all_even :
            continue
        count += 1
    print(count)


if __name__ == '__main__':
    main()
