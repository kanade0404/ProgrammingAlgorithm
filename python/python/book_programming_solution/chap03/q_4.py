import sys
def main():
    a = map(int, input().split())
    max_val, min_val = 0, sys.maxsize
    for i in a:
        if i > max_val:
            max_val = i
        if i < min_val:
            min_val = i
    print(max_val - min_val)

if __name__ == '__main__':
    main()
