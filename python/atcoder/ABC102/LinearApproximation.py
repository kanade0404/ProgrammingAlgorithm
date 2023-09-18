# これにめっちゃ時間かかってしまい他が出来なかった・・・

def processing_min_sorrow(x, b, z):
    return x - (b + z)


if __name__ == '__main__':
    num = int(input())
    line = input().split()
    max = 10000
    count = 0
    result1 = 0
    result2 = 0
    while True:
        x = 0
        for l in line:
            x += int(l)
        z = 0
        for n in range(num):
            z += n
        up_result = processing_min_sorrow(x, count, z)
        down_result = processing_min_sorrow(x, count * -1, z)
        if up_result < max:
            result1 = up_result
        if down_result < max:
            result2 = down_result
        count += 1
    print(final_result)

from statistics import *
n = int(input())
list = input().split()
for i in range(n):
    list[i] = int(list[i]) - (i + 1)
b = int(median(list))
for i in range(n):
    list[i] = abs(list[i] - b)
print(sum(list))


# user:mo_mo_
from statistics import *



n = int(input())
list = input().split()

for i in range(n):
    list[i] = int(list[i]) - (i + 1)

b = int(median(list))

for i in range(n):
    list[i] = abs(list[i] - b)

print(sum(list))

