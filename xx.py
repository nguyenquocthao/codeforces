from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd
from functools import lru_cache


def union(a,b):
    res=[]
    i,j=0,0
    while i<len(a) and j<len(b):
        if a[i]>b[j]:
            res.append(a[i])
            i+=1
        else:
            res.append(b[j])
            j+=1
    return res

a,b = [10,9,3,1], [20,2,1]
print(a,b,union(a,b))