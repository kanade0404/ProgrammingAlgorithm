# result:AC
a1, a2, a3 = map(int, input().split())
x = abs(a1 - a2) + abs(a2 - a3)
y = abs(a1 - a3) + abs(a1 - a2)
z = abs(a2 - a3) + abs(a1 - a3)
print(min(x, y, z))

# user:kyuna
m,_,M=sorted(map(int,input().split()))
print(M-m)
