import redis
import uuid 
import random
from itertools import product, combinations
import string

# pw = "8RMo89JBm6"
# client = redis.StrictRedis(host='localhost', port=6380, password=pw)

# with open("audio.wav", "rb") as f:
#     file_bytes = f.read()
#     id = str(uuid.uuid4())
#     for i in range(2000):
#         client.set("key-audio" + id + str(i), file_bytes, ex=3600)
#         print(i)

# print("Number of keys", client.dbsize())

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




