# My code
# result:AC
lines = input()
count = 0
for line in lines:
    count += int(line)
if int(lines) % count == 0:
    print('Yes')
else:
    print('No')


# Answer example
N=int(input())
print(['Yes','No'][N%sum(map(int,str(N)))>0])
