N = abs(int(input()))
count = abs(N)
result = ''
if N == 0:
    print('0')
else:
    for n in range(N):
        if count <= 1:
            break
        if count % 2 == 0:
            if count / 2 == 1:
                result = result + '0'
            else:
                result = result + '1'
        else:
            result = result + '0'
            count -= 1
        count = int(count / 2)
    print('1' + result)

# user:Shuzaei
a=int(input())
b=''
while a:
 b=str(a&1)+b
 a=-(a//2)
print(b if b else 0)
