from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd
from functools import lru_cache, cmp_to_key

# def convert(v):
#     if v<=3: return str(v)
#     return convert((v-1)//3) + str((v-1)%3 + 1)
    


# n=100000
# q=100000

# def compare(a,b):
#     if a+b<b+a: return -1
#     elif a+b==b+a: return 0
#     else: return 1

# data=[convert(v) for v in range(1, 40)]
# print(data)
# print(sorted(data))
# print(sorted(data, key=cmp_to_key(compare)))

# for i in range(100):
#     print(1<<(10 - i.bit_count()))

l = [1<<(8 - i.bit_count()) for i in range(256)]
print(sorted(l, reverse=True))