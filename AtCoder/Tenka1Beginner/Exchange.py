a, b, K = map(int, input().split())
count = 0
while True:
    if a % 2 == 1:
        a -= 1
    b += a / 2
    a -= a / 2
    count += 1
    if count == K:
        break
    if b % 2 == 1:
        b -= 1
    a += b / 2
    b -= b / 2
    count += 1
    if count == K:
        break
print(int(a), int(b))
