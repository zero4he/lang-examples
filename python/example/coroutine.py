from asyncio.coroutines import coroutine


# 生产
def product(num):
    for x in range(0, num):
        yield x

# 消费
for x in product(10):
    print(x)

