# user:jellyfish
d, n, m = map(int, input().split())

if m != 0:
    print(((2 / pow(d, 2)) * (d - m)) * (n - 1))
else:
    print((1 / d) * (n - 1))
