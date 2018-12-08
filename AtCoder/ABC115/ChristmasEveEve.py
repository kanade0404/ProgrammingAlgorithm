N = int(input())
result = []
for idx in range(N):
    result.append(int(input()))
result.sort(reverse=True)
result[0] = int(result[0] / 2)
print(sum(result))
