count = int(input())
for l in input():
    if count == 0:
        print('Yes')
        break
    if l == '+':
        count += 1
    else:
        count -= 1
if count != 0:
    print('No')

# user: ykst51
a=int(input())
r="No"if a else"Yes"
for i in input():
    a+=1if i=='+'else-1
    if a==0:r="Yes"
print(r)
