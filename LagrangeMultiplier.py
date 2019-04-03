import numpy as np
from scipy.optimize import fsolve


def lagrange(k):
    x, y, lag = k
    return -(x ** 2) - (y ** 2) - lag * (x + y - 1)


def pertial_derivative(x):
    d_lambda = np.zeros(len(x))
    h = 1e-3
    for i in range(len(x)):
        dx = np.zeros(len(x))
        dx[i] = h
        d_lambda[i] = (lagrange(x + dx) - lagrange(x - dx)) / (2 * h)
    return d_lambda


max_arg = fsolve(pertial_derivative, [2, 1, 3])
print('max arguments: x = {0:.2}, y = {1:.2}'.format(max_arg[0], max_arg[1]))
print('max value: {0:.2}'.format(lagrange(max_arg)))
