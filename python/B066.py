def main():
    H, W = map(int, input().split())
    result, count = '', 0
    for _ in range(H):
        degs = [int(i) for i in input().split()]
        for i in range(len(degs)):
            if sum(degs[:i]) == sum(degs[i:]):
                result += 'A' * i + 'B' * (5 - i) + '\n'
                count += 1
    print('Yes\n{}'.format(result) if count == 3 else 'No')


if __name__ == '__main__':
    main()
