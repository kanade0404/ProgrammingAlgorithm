K = int(input())
N = K - 1
count = 0
for i in range(K):
    I = i + 1
    for j in range(N):
        J = j + 1
        if I % 2 != 0 & J % 2 == 0:
            count += 1
        elif I % 2 == 0 & J % 2 != 0:
            count += 1
    N -= 1
print(count)

# user:kyuna
print(int(input())**2>>2)
