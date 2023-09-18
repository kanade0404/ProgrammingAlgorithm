# user:yaketake08
N, M = map(int, input().split())
*A, = map(int, input().split())
C = {0: 1}
ans = 0
s = 0
for a in A:
    s = (s + a) % M
    ans += C.get(s, 0)
    C[s] = C.get(s, 0) + 1
print(ans)
