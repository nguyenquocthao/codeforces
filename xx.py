from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd
from functools import lru_cache, cmp_to_key
from typing import List

def divceil(n,k):
    return (n+k-1)//k

res = set()
def dp(arr, i, k):
    if i<0 or i>=len(arr): return
    if k==0:
        if arr[-1]>0:  res.add(tuple(arr))
        # print(arr)
        return
    arr[i]+=1
    dp(arr[:], i-1, k-1)
    dp(arr[:], i+1,k-1)

dp([0,0,0,0,0,0], 0, 13)
# dp([0,0,0],0,5,True)
print(len(res))
for v in sorted(res):
    print(v)
    
    
# 0 -> 1 -> 2 -> 1 -> 0