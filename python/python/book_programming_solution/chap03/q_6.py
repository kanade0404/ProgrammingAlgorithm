
def main():
    k, n = map(int, input().split())
    a = [int(input())for _ in range(n)]
    result = 0
    for bit in range(2 ** n):
        sum = 0
        for i in range(k):
