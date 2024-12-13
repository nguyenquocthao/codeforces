from itertools import permutations, combinations, product
from collections import defaultdict
import random
from math import gcd
from functools import lru_cache


def canon(s):
    res=s
    for i in range(len(s)):
        s2=s[i:] + s[:i]
        res=min(res, s2)
        # res=min(res, s2, s2[::-1])
    return res

def gen(ncolor,n):
    for l in product(range(ncolor),repeat=n):
        yield ''.join(str(v) for v in l)

data = list(gen(2,7))


scanon = set()
for s in data:
    # print(s, canon(s))
    scanon.add(canon(s))
print(len(scanon), sorted(scanon,key=lambda v: v.count('1')))


def cal(ncolor,n):
    s=0
    for i in range(n):
        # print(i,  gcd(n,i), n//gcd(n, i), pow(ncolor, n//gcd(n, i)))
        s += pow(ncolor, gcd(n, i))
    if n%2==1: s+=n*pow(ncolor, n//2)
    else: s += (n//2) * pow(ncolor, n//2) + (n//2) * pow(ncolor, (n//2) + 1)
    return s//(2*n)
    # return s//n
print(cal(3,8))


# def tran1(s,i):
#     return s[i:] + s[:i]
# def tran2(s,i):
#     if i==0: return s
#     else: return s[::-1]
# for i,j in product(range(6),range(2)):
#     d=[]
#     for s in data:
#         s1 = tran1(s, i)
#         s2 = tran2(s1, j)
#         if s==s2: d.append(s)
#     # print(i,j,len(d), d)



# def run(n):
#     maxv, res = 0, []
#     for l in permutations(list(range(1,n+1))):
#         s=0
#         for i,j in combinations(range(n+1),2): s+= min(l[i:j])
#         if s>maxv: maxv,res = s, [l]
#         elif s==maxv: res.append(l)
#     print(maxv)
#     print(len(res), res)

# run(3)
# run(4)
# run(5)
# run(6)


# for i,j,k in combinations(range(10),3):
#     for x in range(10, 0, -1):
#         if i%x==j%x==k%x: 
#             print(i,j,k,x)
#             break

# a = a0x + y
# b = b0x + y
# y < min(a,b)
# a-b = zz x


# def check(s):
#     n=len(s)
#     data=[1 if ch=='1' else -1 for ch in s]
#     @lru_cache(None)
#     def dp(i, ngroup):
#         if ngroup==0: return 0
#         if i==ngroup: return ngroup*data[i] + dp(i-1,ngroup-1)
#         return ngroup*data[i] + max(dp(i-1,ngroup), dp(i-1,ngroup-1))
#     for i in range(n):
#         print(i, dp(n-1,i))
#     d2,s=[],0
#     for i in range(n-1,0,-1):
#         s+=data[i]
#         d2.append((s,i))
#     d2.sort(reverse=True)
#     print(d2)


# words=['1001', '1010', '0110', '001110', '1111111111', '11111', '000000000011111', '0'*20 + '1'*5, ('00001111'*3)]
# # check()
# for w in words:
#     print(w)
#     check(w)

