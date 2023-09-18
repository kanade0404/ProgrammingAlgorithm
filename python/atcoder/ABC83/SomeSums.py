N, A, B = map(int, input().split())
print(sum(i * (A <= sum(map(int, str(i))) <= B) for i in range(N + 1)))
