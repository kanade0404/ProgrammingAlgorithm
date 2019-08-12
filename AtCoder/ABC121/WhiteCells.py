H, W = map(int, input().split())
l = list()
for _ in range(H):
    l.append([i for i in range(W)])
h, w = map(int, input().split())
print(l[0])
print(l[1])
print(l[2])
for i in range(h):
    print(i)
    l[h] = [0 for _ in range(W)]
    print(l)


