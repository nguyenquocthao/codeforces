from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd
from functools import lru_cache, cmp_to_key
from typing import List

def divceil(n,k):
    return (n+k-1)//k

# @lru_cache(None)
# def stirlingsecond(n,k):
#     if n==0 and k==0: return 1
#     if n==0 or k==0: return 0
#     return k * stirlingsecond(n-1,k) + stirlingsecond(n-1,k-1)

# for i,j in product(range(15),repeat=2):
#     if stirlingsecond(i,j)%2==1: print(i,j, i-divceil(j+1,2), (j-1)//2)
#     # print(i,j, stirlingsecond(i,j))
    

def xorupto(n):
    return [n,1,n+1,0][n%4]
    # if n%4==0: return 

def xorif(mask, m):
    res=0
    for i in range(1,m+1):
        if i|mask==i: res^=i
    return res

def xorif2(mask:int,m:int):
    if m<mask: return 0
    x = m-mask
    bb = []
    for b in range(m.bit_length()):
        if (1<<b) & mask > 0: 
            if (1<<b) & x > 0: x|=(1<<b)-1
            continue
        bb.append(b)
    res=0
    for b,ii in enumerate(bb):
        if (1<<ii) & x>0: res|=1<<b
        # if (1<<b) & x > 0: res|=1<<nb
        # nb+=1
    # res+=1
    # print("xorif2", mask, m, x, res)
    # also have 0
    base = 0 if res%2==1 else mask
    val = xorupto(res)
    for b,ii in enumerate(bb):
        if val &(1<<b) > 0: base|=1<<ii
    return base
    
    
        




    


def count(mask, maxv):
    return sum(i|mask==i for i in range(maxv+1))

# def count2(mask, maxv):
#     if mask > maxv:
#         return 0  # If `mask` itself is larger than `maxv`, no valid `i` exists.

#     # Free bits are those not set in `mask`, but within the range of `maxv`
#     free_bits = ~mask & maxv
#     count = 0

#     # Iterate over all subsets of `free_bits`
#     subset = 0
#     while True:
#         i = mask | subset  # Valid number `i`
#         if i > maxv:
#             break
#         count += 1

#         # Generate the next subset of `free_bits`
#         if subset == free_bits:
#             break
#         subset = (subset + 1) & free_bits

#     return count

for mask, m in product(range(20),repeat=2):
    mask+=1
    if mask>m: continue
    # xorif2(mask, m)
    # print(count(mask,m))
    print(mask,m, xorif(mask,m),xorif2(mask,m))
    
