import numpy as np
H, W = map(int, input().split())
list1 = []
for h in range(H):
    R = input()
    l = []
    for r in R:
        l.append(r)
    list1.append(l)
list2 = []
for l in list1:
    if l.count(".") != W:
        list2.append(l)
list3 = []

