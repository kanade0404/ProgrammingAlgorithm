def func(n: int) ->int:
    if n == 0:
        return 0
    return n + func(n - 1)
def main():
    n = int(input())
    print(func(n))
if __name__ == '__main__':
    main()
