N, T = map(int, input().split())
result_cost = 0
for n in range(N):
    c, t = map(int, input().split())
    if t <= T:
        if result_cost == 0:
            result_cost = c
        if c < result_cost:
            result_cost = c
print('TLE' if result_cost == 0 else result_cost)
