from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd
from functools import lru_cache

def check(s):
    n=len(s)
    data=[1 if ch=='1' else -1 for ch in s]
    @lru_cache(None)
    def dp(i, ngroup):
        if ngroup==0: return 0
        if i==ngroup: return ngroup*data[i] + dp(i-1,ngroup-1)
        return ngroup*data[i] + max(dp(i-1,ngroup), dp(i-1,ngroup-1))
    for i in range(n):
        print(i, dp(n-1,i))
    d2,s=[],0
    for i in range(n-1,0,-1):
        s+=data[i]
        d2.append((s,i))
    d2.sort(reverse=True)
    print(d2)


words=['1001', '1010', '0110', '001110', '1111111111', '11111', '000000000011111', '0'*20 + '1'*5, ('00001111'*3)]
# check()
for w in words:
    print(w)
    check(w)

