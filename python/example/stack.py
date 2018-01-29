stack = ["A", "B", "C", "D"]

del stack[2]

stack.append("E")
stack.append("F")

print(stack)

i = stack.index("D")

# i = stack.index("C", 2)  # ValueError: 'C' is not in list
i = stack.index("D", 2)
print(i)

try:
    i = stack.index("C", 2)
except Exception as e:
    print(Exception, ":", e)
