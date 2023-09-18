# My code
# result:WA
n, k = map(int, input().split())
a = input().split()
line = []
for i in a:
    line.append((int(i)))
print(line)
count = 0
# 1の両側をdeleteして最後1だけ残る方法で考えたけど上手くいかず
while line.count(1) != 0:
    idx = line.index(1)
    try:
        line.pop(idx - 1)
    except:
        pass
    try:
        line.pop(idx + 1)
    except:
        pass
    finally:
        count += 1
print(count)

# Answer example
# user:makk
N,K=map(int,input().split(' '))
print( (N-1+K-2) // (K-1))
