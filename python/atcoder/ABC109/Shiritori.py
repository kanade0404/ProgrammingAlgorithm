N = int(input())
list = []
result = True
for i in range(N):
    word = input()
    if i > 0:
        chk_word = list[i - 1]
        if chk_word[-1] != word[0]:
            result = False
    list.append(word)
    if list.count(word) > 1:
        result = False
print('Yes' if result else 'No')
