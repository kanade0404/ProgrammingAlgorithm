N = int(input())
check = False
if N % 4 == 0 or N % 7 == 0:
    check = True
else:
    for i in range(25):
        for j in range(25):
            if 4 * (i + 1) + 7 * (j + 1) == N:
                check = True
                break
if check is True:
    print('Yes')
else:
    print('No')

# user:ransewhale
print("No" if int(input()) in [1,2,3,5,6,9,10,13,17] else "Yes")
