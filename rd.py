import redis
import uuid 
import random
from itertools import product, combinations
import string

with open('x.txt','w') as f:
    n,m = 2000000, 1000000
    print(n,m, file=f)
    a,b = ['-']*n, ['-']*m
    pool=string.ascii_lowercase[:3] + '-*'
    for i in range(n):
        j = random.randint(0, len(pool)-2)
        a[i]=pool[j]
    for i in range(m):
        j=random.randint(0, len(pool)-2)
        b[i]=pool[j]
    b[0]='*'
    b[-1]='*'
    print(''.join(a), file=f)
    print(''.join(b), file=f)




