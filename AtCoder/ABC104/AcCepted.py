import re
s = input()
print('AC' if s[0].isupper() and (s[2:-2].count('C') == 1 or (s[2:-2] == '' and s[2] == 'C')) and len(re.findall('[a-z]', s)) == len(s) - 2 else 'WA')

# c = 'C'
# if s[0].isupper():
#     S = s[2:-2]
#     if S.count('C') == 1 or (S == '' and s[2] == 'C'):
#         if len(re.findall('[a-z]', s)) == len(s) - 2:
#             print('AC')
