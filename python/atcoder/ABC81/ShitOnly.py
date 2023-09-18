_ = int(input())
isEven = True
l1 = map(int, input().split())
count = 0
l2 = []
while isEven:
    for i in l1:
        if i % 2 == 0:
            l2.append(i / 2)
        else:
            isEven = False
            break
    if isEven:
        l1 = l2.copy()
        l2 = []
        count = count + 1
print(count)
