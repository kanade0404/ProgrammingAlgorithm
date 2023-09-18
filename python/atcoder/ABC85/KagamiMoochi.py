N = int(input())
d = [int(input()) for i in range(N)]
d.sort(reverse=True)
print(len([i for i in range(N - 1) if d[i] - d[i + 1] > 0]) + 1)
