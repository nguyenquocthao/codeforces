from itertools import permutations, combinations, product
from collections import defaultdict,Counter
import random
from math import gcd, lcm
from functools import lru_cache, cmp_to_key
from typing import List

# def cal(x,y):
#     y=1<<y
#     return sum(i^y < i for i in range(x+1))

# def nbitat(x, i):
#     p = 1<<i
#     a,b = divmod(x, 2*p)
#     return a*p + max(0, b-p+1)

# # for x,y in product(range(1,20),(1,2,4,8,16)):
# #     print(x,y, cal(x,y))

# for y,x in product(range(5),range(1,20)):
#     print(x,y,cal(x,y), nbitat(x,y))

data=[1,1]
while len(data)<100:
    data.append(data[-1]+data[-2])

print(data)
for target in range(2,21):
    print(target, [i+1 for i,v in enumerate(data) if v%target==0])