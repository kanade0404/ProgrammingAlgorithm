def from_top(array, N, M, count):
    for i in range(N - 1):
        for j in range(M - 1):
            if array[i + 1, j] == '#':
                break
            else:
                count += 1


def from_right(array, N, M, count):
    for i in range(N - 1):
        for j in range(M - 1):
            if array[i + 2, j] == '#':
                break
            else:
                count += 1


def from_bottom(array, N, M, count):
    for i in range(N - 1):
        for j in range(M - 1):
            if array[2 - i, j] == '#':
                break
            else:
                count += 1


def from_left(array, N, M, count):
    for i in range(N - 1):
        for j in range(M - 1):
            if array[i, j+ 1] == '#':
                break
            else:
                count += 1


if __name__ == '__main__':
    N, M = map(int, input().split())
    array = []
    for i in range(N):
        arr = []
        for j in input():
            arr.append(j)
        array.append(arr)
    print(array)
    count = 0
    from_top(array, N, M, count)
    from_right(array, N, M, count)
    from_bottom(array, N, M, count)
    from_left(array, N, M, count)
    print(count)
