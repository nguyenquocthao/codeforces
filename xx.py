from itertools import permutations, combinations, product
from collections import defaultdict,Counter
import random
from math import gcd, lcm
from functools import lru_cache, cmp_to_key
from typing import List

secret = [random.randint(0,9) for _ in range(6)]

pool = [list(x) for x in product(range(10),repeat=6)]
# random.shuffle(pool)

# def match(x,y):
#     bulls = sum(a==b for a,b in zip(x,y))
#     cows = sum((Counter(x)&Counter(y)).values()) - bulls
#     return (bulls, cows)

# def run(x):
#     global pool
#     result = match(secret, x)
#     print("Guessed", x, "result", result)
#     pool = [v for v in pool if match(v, x) == result]
#     # random.shuffle(pool)
#     print(len(pool), pool[:10])

# while len(pool)>1:
#     x = pool[0]
#     run(x)

# print("pool", pool)
# print("secret", secret)

data=[0]*200
for i in range(1, len(data)):
    data[i]=1+data[i//4]
# print(data)
for i,v in enumerate(data):
    print(i,v,i.bit_length())