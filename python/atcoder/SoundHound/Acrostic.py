line = input()
n = int(input())
result = ''
idx = 0
while idx < len(line):
    result = result + line[idx]
    idx = idx + n
print(result)

# user:kyuna
print(input()[::int(input())])
