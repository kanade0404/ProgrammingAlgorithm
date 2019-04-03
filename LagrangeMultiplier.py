import numpy as np
from scipy.optimize import fsolve


def lagrange(X):
    x, y, L = X
    return -(x ** 2) - (y ** 2) - L * (x + y - 1)


def pertial_derivative(X):
    dLambda = np.zeros(len(X))
    h = 1e-3
    for i in range(len(X)):
        dX = np.zeros(len(X))
        dX[i] = h
        dLambda[i] = (lagrange(X + dX) - lagrange(X - dX)) / (2 * h)
    return dLambda


max_arg = fsolve(pertial_derivative, [2, 1, 3])
print('max arguments: x = {0:.2}, y = {1:.2}'.format(max_arg[0], max_arg[1]))
print('max value: {0:.2}'.format(lagrange(max_arg)))
