ints = [0, 1, 2, 3]  # list 是内置的名称, 不用使用list作为变量名

# ints.append(2)

print(ints[1:3])  # 左闭右开区间,输出: [1, 2]
print(ints[1:-1])  # 左闭右开区间,输出:[1, 2]

del ints[1]
print(1 in ints)  # False
print(3 in ints)  # True

ints.append(9)

print(max(ints))

# 元组转list
aTuple = (3, 4)
l = list(aTuple)
print(l)

# filter
f1 = filter(lambda x: x > 2, [1, 2, 3, 4])
f2 = [x for x in [1, 2, 3, 4] if x > 2]
print(type(f1))  # <class 'filter'>  filter可迭代
print("f1: ", list(f1), "     f2: ", f2)  # f1:  [3, 4]      f2:  [3, 4]

# map
m1 = map(lambda x: x + 10, [1, 2, 3, 4])
m2 = [x + 10 for x in [1, 2, 3, 4]]
print("m1: ", list(m1), "     m2: ", m2)  # m1:  [11, 12, 13, 14]      m2:  [11, 12, 13, 14]
