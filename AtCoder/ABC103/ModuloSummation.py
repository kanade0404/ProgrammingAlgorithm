# user:AC
N = int(input())
line = input().split()
r = 0
for i in range(N):
    r += int(line[i]) - 1
print(r)

# user:tnktemp
print(-int(input())+sum(map(int,input().split())))
