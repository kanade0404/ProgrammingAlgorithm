H, A = map(int, input().split())
print(int(H / A)) if H % A == 0 else print(int(H / A) + 1)
