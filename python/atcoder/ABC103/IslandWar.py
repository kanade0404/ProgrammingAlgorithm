N, M = map(int, input().split())
a = []
b = []
for i in range(M):
    a[i], b[i] = map(int, input().split())


# user:kenkoooo
nm = [int(x) for x in input().split()]
N = nm[0]
M = nm[1]

to = [N] * N
for _ in range(M):
    nm = [int(x) for x in input().split()]
    A = nm[0] - 1
    B = nm[1] - 1
    if to[B] == N or to[B] < A:
        to[B] = A

last = 0
ans = 0
for i in range(N):
    if to[i] == N:
        continue
    if to[i] >= last:
        last = i
        ans += 1
print(ans)
