# My code
# result:AC
lines = input()
count = 0
for line in lines:
    if line == '+':
        count += 1
    elif line == '-':
        count -= 1
print(count)


# Answer example
# user:morio_
print(2*input().count('+')-4)
