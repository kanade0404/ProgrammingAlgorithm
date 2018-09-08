import math
from functools import reduce
_, X, *A = map(int, open(0).read().split())
print(reduce(math.gcd, [abs(a - X) for a in A]))
