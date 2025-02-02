from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd
from functools import lru_cache, cmp_to_key


# def run(l):
#     print(l)
#     n=len(l)
#     for target in range(1,11):
#         for i in range(n):
#             nlo, nhi=0,0
#             for j in range(i,n):
#                 if l[j]<target: nlo+=1
#                 elif l[j]>target: nhi+=1
#                 if (j-i)%2==1 and max(nlo,nhi)*2 < j-i: print(i,j,target, l[i:j+1])


# run([6, 3, 2, 3, 5, 3, 4, 2, 3, 5])

def diff(a,b):
    return sum(abs(ord(x) - ord(y)) for x,y in zip(a,b))


print(diff("abcdefghijijikndndfsa", "bbbeeehhhjjjjjnnndddd"))