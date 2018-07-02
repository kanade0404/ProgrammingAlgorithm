n = int(input())
line = input().split()
result = 0
for i in range(n):
    for j in range(n):
        r = abs(int(line[i]) - int(line[j]))
        if r > result:
            result = r
print(result)

# user:kyuna
input()
A=list(map(int,input().split()))
print(max(A)-min(A))
