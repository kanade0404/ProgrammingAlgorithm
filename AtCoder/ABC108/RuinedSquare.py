x1, y1, x2, y2 = map(int, input().split())
diffX = x2 - x1
diffY = y2 - y1
x3 = 0
y3 = 0
x4 = 0
y4 = 0
if diffX >= 0 & diffY >= 0:
    x3 = x2 + diffY * -1
    y3 = y2 + diffX
    x4 = x3 + diffX * -1
    y4 = y3 + diffY * -1
elif diffX >= 0 & diffY <= 0:
    x3 = x2 + diffY * -1
    y3 = y2 + diffX
    x4 = x3 + diffX
    y4 = y3 + diffY
elif diffX <= 0 & diffY >= 0:
    x3 = x2 + diffY * -1
    y3 = y2 + diffX
    x4 = x3 + diffX * -1
    y4 = y3 + diffY * -1
elif diffX <= 0 & diffY <= 0:
    x3 = x2 + diffY
    y3 = y2 + diffX * -1
    x4 = x3 + diffX
    y4 = y3 + diffY
print(x3, y3, x4, y4)

# user:oyassan
a,b,c,d = map(int,input().split())

print(b+c-d,c+d-a,b-d+a,c+b-a)
