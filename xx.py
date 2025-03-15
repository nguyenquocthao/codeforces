from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd, lcm
from functools import lru_cache, cmp_to_key
from typing import List

def run(a,b,c):
    base = [a,b,c]
    data=[a,b,c]
    res=max(base)
    for _ in range(100):
        diff = max(data)-min(data)
        print(data, diff)
        i = min(range(3), key=lambda ind:data[ind])
        data[i]+=2*base[i]
        res = min(res, diff)
    return res

run(6,10,15)
        
        
    

def sumsq(a, step, n):
    res=0
    for _ in range(n):
        res+=a*a
        a+=step
    return res
