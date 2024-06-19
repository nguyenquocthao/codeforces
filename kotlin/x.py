import bisect
from itertools import product

cached1=[2]
for i in range(1, 20):
    cached1.append(cached1[-1]+pow(3,i))
print(cached1)

def solve1(n):
    i = bisect.bisect_right(cached1,n)-1
    return cached1[i]

# print(solve1(15))
for i in range(2,30):
    print("solve1", i, solve1(i))


def solve2(data):
    m,n =len(data), len(data[0])
    marked=[[0]*n for _ in range(m)]
    islands = [0]
    def nb(i,j):
        for i2,j2 in ((i+1,j),(i-1,j),(i,j+1),(i,j-1)):
            if 0<=i2<m and 0<=j2<n: yield i2,j2
    def markisland(i,j):
        if marked[i][j]>0: return 0
        marked[i][j] = len(islands)
        res=1
        for i2,j2 in nb(i,j):
            if data[i2][j2]==1: res += markisland(i2,j2)
        return res
    for i,j in product(range(m),range(n)):
        if data[i][j]==1 and marked[i][j]==0:
            islands.append(markisland(i,j))
    res=0
    print(islands, marked)
    for i,j in product(range(m),range(n)):
        if data[i][j]==0:
            nbislands = set(marked[i2][j2] for i2,j2 in nb(i,j))
            res = max(res, 1 +sum(islands[ind] for ind in nbislands))
    return res

data2 = [[1,0,1,0,1,1],
         [1,0,1,0,1,1],
         [1,1,1,0,0,0],
         [0,0,0,1,1,1],
         [1,1,0,0,1,0],
         [1,1,1,0,1,1]]
print(solve2(data2))

cached3=[0]*15
for i in range(1, 15):
    # there are 9 1-digit numbers, 90 2-digit numbers, 900 3-digit numbers
    cached3[i] = cached3[i-1] + 9*i*pow(10,i-1)
print(cached3)

def solve3(k):
    i=bisect.bisect_left(cached3, k) - 1
    k-=cached3[i]
    k-=1
    start,offset = divmod(k, i+1)
    start += pow(10, i)
    # print(start, offset)
    return (str(start)+str(start+1)+str(start+2)+str(start+3))[offset:offset+4]


for i in range(1,30):
    print("solve3", i, solve3(i))
