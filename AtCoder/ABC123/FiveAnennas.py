A = [int(input()) for _ in range(5)]
k = int(input())
for i in range(5):
    for j in range(5):
        print(abs(A[i] - A[j]))
        if abs(A[i] - A[j]) > k:
            print(':(')
            exit()
print('Yay!')
