from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd
from functools import lru_cache, cmp_to_key

# def convert(v):
#     if v<=3: return str(v)
#     return convert((v-1)//3) + str((v-1)%3 + 1)
    


# def run(l,r):
#     print('running', l, r)
#     cur,res = 0, []
#     for a,b,c in combinations(range(l,r+1),3):
#         v = (a^b)+(b^c)+ (c^a)
#         if v>cur: cur,res = v, [(a,b,c)]
#         elif v==cur: res.append((a,b,c))
#     print(cur, res)

# run(0,2)
# run(0,8)
# run(1,3)
# run(6,22)
# run(128,137)

def update(tup, i):
    l=list(tup)
    v, l[i] = divmod(tup[i], 2)
    if i>0: l[i-1]+=v
    if i+1 < len(l): l[i+1]+=v
    return tuple(l)

cached = {}
def run(tup):
    if tup in cached: return max(cached[tup])
    if all(v<=1 for v in tup): 
        cached[tup] = [0]
        return 0
    res=[]
    if tup[0]>=2:
        res.append(tup[0]//2 + run(update(tup,0)))
    for i in range(1, len(tup)):
        if tup[i]>=2:
            res.append(run(update(tup, i)))
    cached[tup] = res
    return max(cached[tup])

    
run((1,3,5,7))
for k in sorted(cached.keys(), key=lambda v: -sum(v)):
    print(k, cached[k])
