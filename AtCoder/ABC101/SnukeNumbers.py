# My code
# not submit
# ほとんどできなかった
# と言うかすぬけ數の意味が理解できなかったので読解力（）
k = input()
n = 1
for i in k:
    n += int(i)


# Answer example
# user:yamadah
K = int(input())

def S(n):
    return sum(map(int, list(str(n))))

def norm(n):
    return n/S(n)

snk = 1
order = 0
for i in range(K):
    print(snk)
    nxt1 = snk + 10 ** order
    nxt2 = snk + 10 ** (order + 1)
    if norm(nxt2) < norm(nxt1):
        order += 1
        snk = nxt2
    else:
        snk = nxt1
