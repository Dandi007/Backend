"""2020.05.11,展示使用dataclass来自动生成多field class的大小比较的函数
"""
from dataclasses import dataclass, field


@dataclass(order=True)
class A:
    first: int
    second: int
    third: int = field(compare=False)


a = A(1, 2, 3)
b = A(1, 2, 2)
c = A(2, 1, 3)
d = A(0, 3, 1)

# 第三个field不参与比较
print(a == b)

# 先比较第一个filed
print(a > c)
print(c > a)
print(a > d)
