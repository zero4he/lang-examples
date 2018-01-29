from functools import reduce

# lambda是一个表达式，函数体比def简单很多,同时lambda也是一个function
f = lambda x, y: x + y

print(f(2, 3))


def fun(x):
    return lambda y: x + y


# fun(9) 返回的是一个lambda, 这个lambda可以接收一个参数并返回一个值
print(fun(9)(10))

# 下面的两句等同: print(fun(9)(10))
f = fun(9)
print(f(10))

v = reduce(lambda x, y: x * y, range(1, 5))

print(v)
