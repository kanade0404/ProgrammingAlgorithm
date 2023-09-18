a, b = map(int, input().split())
if a + b == 15:
    print('+')
elif a * b == 15:
    print('*')
else:
    print('x')

# user: funwarioisii
a, b = map(int, input().split())
print('+' if a + b ==15 else '*' if a * b == 15 else 'x')
