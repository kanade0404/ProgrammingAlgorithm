N = int(input())
A = []
for n in range(N):
    A.append(int(input()))
A.sort()
result = []
diff = 0
if len(A) % 2 == 0:
    diff = A[-1] - A[0]
    for a in range(len(A) - 1):
        result.append(A[a + 1])
else:
    result = A
for n in range(N - 1):
    if n % 2 == 0:
        diff += abs(result[-n - 1] - result[n])
    else:
        diff += abs(result[-n] - result[n])
print(diff)
