N, K = map(int, input().split())
result = 0
for i in range(N):
    N1 = N
    for v in range(N1):
        N2 = N
        for r in range(N2):
            if (((i + 1) + (v + 1)) % K == 0) & (((i + 1) + (r + 1)) % K == 0) & (((v + 1) + (r + 1)) % K == 0):
                result += 1
        N2 -= 1
    N1 -= 1
print(result)

# user:toku
n, k = map(int, input().split())

if k % 2 == 1:
    print((n//k)**3)
else:
    print((n//k)**3+((n//(k//2)+1)//2)**3)
