N, K = map(int, input().split())
tree = [int(input()) for n in range(N)]
diff = 10 ** 9
tree.sort()
for i in range(N - (K - 1)):
    if tree[i] == tree[i + 1] == tree[i + 2]:
        diff = 0
        break
    elif tree[i + K - 1] - tree[i] < diff:
        diff = tree[i + K - 1] - tree[i]
print(diff)
