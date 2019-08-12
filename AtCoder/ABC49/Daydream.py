s = input()
flg = True
while(flg):
    if s.endswith('dream') or s.endswith('erase'):
        s = s[:-5]
    elif s.endswith('eraser'):
        s = s[:-6]
    elif s.endswith('dreamer'):
        s = s[:-7]
    elif len(s) == 0:
        flg = False
        print("YES")
    else:
        flg = False
        print('NO')
