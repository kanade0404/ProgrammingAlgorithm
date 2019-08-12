K, X = map(int, input().split())
print(' '.join(map(str, [i for i in range(X - K + 1, X + K)])))
