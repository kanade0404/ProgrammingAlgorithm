def gcd(m: int,n: int)->int:
    if n== 0:
        return m
    return gcd(n, m%n)

def main():
    m, n = map(int, input().split())
    print(gcd(m, n))

if __name__ == '__main__':
    main()
