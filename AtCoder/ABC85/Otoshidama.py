N, Y = map(int, input().split())
y = [0, 0, 0]
if Y / N != 0:
    print('-1 -1 -1')

for i in range(N):
    if Y - 10000 >= 0:
        y[0] += 1
        Y -= 10000
    elif Y - 5000 >= 0:
        y[1] += 1
        Y -= 5000
    else:
        y[2] += 1
        Y -= 1000
if 10000 * N < Y:
    print('-1 -1 -1')
else:
    print(y[0], y[1], y[2])
