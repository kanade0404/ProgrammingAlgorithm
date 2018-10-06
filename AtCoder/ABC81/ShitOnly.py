N = int(input())
isEven = True
list = []
for i in input().split():
    if int(i) % 2 == 1:
        isEven = False
    list.append(int(i))
