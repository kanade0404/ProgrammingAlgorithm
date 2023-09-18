# user:morio_
import math

L = int(input())
r = int(math.log2(L))

# r+1頂点
vertex = r + 1
edge = 2 * r
res = []
for i in range(1,r+1):
    res.append([i,i+1,0])
    res.append([i,i+1,2**(i-1)])

for i in range(r,0,-1):
    if L-2**(i-1)>=2**r:
        L -= 2**(i-1)
        res.append([i,vertex,L])
        edge += 1

print(vertex,edge)
for l in res:
    print(*l)
