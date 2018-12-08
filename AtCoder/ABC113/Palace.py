N = input()
T, A = map(int, input().split())
H = [abs(T-0.006*x-A) for x in map(int, input().split())]
print(H.index(min(H))+1)
