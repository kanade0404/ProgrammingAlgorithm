# user:AC
X = input()
Y = input()
for i in range(len(Y)):
    if X == Y:
        print('Yes')
        break
    else:
        Y = Y[1:] + Y[0]
    if i == len(Y) - 1:
        print('No')
# user:kyuna
S,T=input()*2,input()
print(['No','Yes'][T in S])
