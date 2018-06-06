# 6種類のアルファベット "u,d,c,s,t,b" から成るn文字の文字列を全て列挙するプログラムを作成してください。
# ただし、nは1以上の整数とし、標準入力より指定されます。その他の条件は(1)と同じであるものとします。
import itertools
lines = ['u', 'd', 'c', 's', 't', 'b']
# 任意の整数
n = int(input())
for i in itertools.product(lines, repeat=n):
    result = ''
    for j in range(n):
        result = result + i[j]
    if 'uud' in result:
        print(result)
print(len(list(itertools.product(lines, repeat=n))))
