from sortedcontainers import SortedList
from typing import List
from collections import deque, Counter, defaultdict
from itertools import combinations, product, permutations
from functools import lru_cache

class RMQ:
    def __init__(self, l):
        n=len(l)
        def minat(i,j):
            if l[i]<=l[j]: return i
            else: return j
        data=[list(range(n))]
        for p in range(n.bit_length()):
            mul,x,prev = 1<<p, [], data[-1]
            for i in range(len(prev)-mul):
                x.append(minat(prev[i], prev[i+mul]))
            data.append(x)
        def query(i,j):
            if i>j: i,j=j,i
            if i==j: return i
            nb = (j-i+1).bit_length()-1
            return minat(data[nb][i], data[nb][j+1-(1<<nb)])
        self.query=query

class LCA:
    def __init__(self, mchild):
        order,depth=[],[]
        def dp(i,d):
            order.append(i)
            depth.append(d)
            for j in mchild[i]:
                dp(j,d+1)
                order.append(i)
                depth.append(d)
        dp(0,0)
        rvindex=[-1]*len(mchild)
        for i,node in enumerate(order): rvindex[node]=i
        if len(order)<16:
            rmq = RMQ(depth)
            def query(a,b):
                i,j = rvindex[a], rvindex[b]
                return order[rmq.query(i,j)]
            self.query=query
            return
        bsize=len(depth).bit_length()//2
        while len(depth)%bsize>0: depth.append(depth[-1]+1)
        blocks=[]
        mblock={}
        for i in range(0, len(depth), bsize):
            key=0
            for j in range(i, i+bsize-1):
                key<<=1
                if depth[j]<depth[j+1]:key|=1
            if key not in mblock:
                mblock[key]=RMQ(depth[i:i+bsize])
            blocks.append(mblock[key])
        def bquery(bi, i, j):
            ind = blocks[bi].query(i,j)
            return ind + bi*bsize
        def getmin(*indexes):
            i=min(indexes, key=lambda j: depth[j])
            return order[i]
        rmqall = RMQ([depth[bquery(ind,0,bsize-1)] for ind in range(len(blocks))])
        def query(a,b):
            i,j = rvindex[a], rvindex[b]
            if i>j: i,j=j,i
            bi,bj = i//bsize, j//bsize
            if bi==bj:
                return getmin(bquery(bi, i-bi*bsize, j-bi*bsize))
            indexes=[bquery(bi, i-bi*bsize, bsize-1), bquery(bj, 0, j-bj*bsize)]
            if bj-bi>=2:
                bk = rmqall.query(bi+1,bj-1)
                indexes.append(bquery(bk, 0, bsize-1))
            return getmin(*indexes)
        self.query=query

class Trie:
    def __init__(self):
        self.c=defaultdict(Trie)
        self.i = 0
        self.end=False
        self.prev=None
    def add(self, w):
        cur=self
        for i,ch in enumerate(w):
            cur=cur.c[ch]
            cur.i=i+1
        cur.end=True
    def next(self, ch):
        if ch in self.c: return self.c[ch]
        if not self.prev: return self
        return self.prev.next(ch)
    def aho_corasick(self):
        dq=deque([(self, y) for y in self.c.values()])
        while dq:
            x,y = dq.popleft()
            y.prev=x
            for ch in y.c:
                dq.append((x.next(ch), y.c[ch]))

def sort_cyclic_shifts(s):
    n = len(s)
    alphabet = 256  # Number of possible characters (ASCII size)
    
    p = [0] * n  # Sorted order of suffixes
    c = [0] * n  # Equivalence class for cyclic shifts

    # Count the frequency of each character
    cnt = [0] * max(alphabet, n)
    for i in range(n):
        cnt[ord(s[i])] += 1
    for i in range(1, alphabet):
        cnt[i] += cnt[i - 1]
    for i in range(n - 1, -1, -1):
        p[cnt[ord(s[i])] - 1] = i
        cnt[ord(s[i])] -= 1

    # Assign the equivalence class for each character
    c[p[0]] = 0
    classes = 1
    for i in range(1, n):
        if s[p[i]] != s[p[i - 1]]:
            classes += 1
        c[p[i]] = classes - 1

    # k is the current length of cyclic shifts being considered
    pn = [0] * n
    cn = [0] * n
    k = 0
    while (1 << k) < n:
        for i in range(n):
            pn[i] = (p[i] - (1 << k) + n) % n
        
        # Radix sort based on the second half of the pair
        cnt = [0] * classes
        for i in range(n):
            cnt[c[pn[i]]] += 1
        for i in range(1, classes):
            cnt[i] += cnt[i - 1]
        for i in range(n - 1, -1, -1):
            p[cnt[c[pn[i]]] - 1] = pn[i]
            cnt[c[pn[i]]] -= 1

        # Recalculate equivalence classes based on new sorted order
        cn[p[0]] = 0
        classes = 1
        for i in range(1, n):
            curr = (c[p[i]], c[(p[i] + (1 << k)) % n])
            prev = (c[p[i - 1]], c[(p[i - 1] + (1 << k)) % n])
            if curr != prev:
                classes += 1
            cn[p[i]] = classes - 1
        c, cn = cn, c
        k += 1

    return p


def build_suffix_array(s):
    # Append a sentinel character '$' which is lexicographically smaller than any character in the string
    # s += '$'
    # return sort_cyclic_shifts(s)[1:]  # Remove the first entry because it corresponds to the '$'
    if isinstance(s, str):
        s = [ord(ch) for ch in s]
    return sa_is(s)


def sa_is(s):
    # Shift all values to be non-negative
    min_val = min(s)
    shifted_s = [x - min_val for x in s]
    upper = max(shifted_s)
    n = len(shifted_s)

    if n == 0:
        return []
    if n == 1:
        return [0]
    if n == 2:
        return [0, 1] if shifted_s[0] < shifted_s[1] else [1, 0]

    sa = [-1] * n
    ls = [False] * n  # Type array
    for i in range(n - 2, -1, -1):
        ls[i] = shifted_s[i] < shifted_s[i + 1] or (shifted_s[i] == shifted_s[i + 1] and ls[i + 1])

    sum_l = [0] * (upper + 1)
    sum_s = [0] * (upper + 1)
    for i in range(n):
        if not ls[i]:
            sum_s[shifted_s[i]] += 1
        else:
            sum_l[shifted_s[i] + 1] += 1
    for i in range(upper + 1):
        sum_s[i] += sum_l[i]
        if i < upper:
            sum_l[i + 1] += sum_s[i]

    def induce(lms):
        nonlocal sa
        sa = [-1] * n
        buf = sum_s[:]
        for d in lms:
            if d == n:
                continue
            sa[buf[shifted_s[d]]] = d
            buf[shifted_s[d]] += 1
        buf = sum_l[:]
        sa[buf[shifted_s[n - 1]]] = n - 1
        buf[shifted_s[n - 1]] += 1
        for i in range(n):
            v = sa[i]
            if v >= 1 and not ls[v - 1]:
                sa[buf[shifted_s[v - 1]]] = v - 1
                buf[shifted_s[v - 1]] += 1
        buf = sum_l[:]
        for i in range(n - 1, -1, -1):
            v = sa[i]
            if v >= 1 and ls[v - 1]:
                buf[shifted_s[v - 1] + 1] -= 1
                sa[buf[shifted_s[v - 1] + 1]] = v - 1

    lms_map = [-1] * (n + 1)
    m = 0
    for i in range(1, n):
        if not ls[i - 1] and ls[i]:
            lms_map[i] = m
            m += 1
    lms = [i for i in range(1, n) if not ls[i - 1] and ls[i]]

    induce(lms)

    if m:
        sorted_lms = [x for x in sa if lms_map[x] != -1]
        rec_s = [0] * m
        rec_upper = 0
        rec_s[lms_map[sorted_lms[0]]] = 0
        for i in range(1, m):
            l = sorted_lms[i - 1]
            r = sorted_lms[i]
            end_l = lms[lms_map[l] + 1] if lms_map[l] + 1 < m else n
            end_r = lms[lms_map[r] + 1] if lms_map[r] + 1 < m else n
            same = True
            if end_l - l != end_r - r:
                same = False
            else:
                while l < end_l:
                    if shifted_s[l] != shifted_s[r]:
                        break
                    l += 1
                    r += 1
                if l != end_l:
                    same = False
            if not same:
                rec_upper += 1
            rec_s[lms_map[sorted_lms[i]]] = rec_upper

        rec_sa = sa_is(rec_s)

        for i in range(m):
            sorted_lms[i] = lms[rec_sa[i]]
        induce(sorted_lms)

    return sa

class Solution:
    def numberOfAlternatingGroups(self, colors: List[int], queries: List[List[int]]) -> List[int]:
        sl=SortedList()
        n=len(colors)
        def get_color_at(i):
            if i<0: return colors[i+n]
            if i>=n: return colors[i-n]
            return colors[i]
        for i in range(n):
            if get_color_at(i)==get_color_at(i+1): sl.add(i)
        def getsl(ind):
            if ind<0: return sl[ind + len(sl)] - n
            if ind>=len(sl): return sl[ind - len(sl)]+n
            return sl[ind]
        count=Counter()
        def inc(x):
            count[x]+=1
        def dec(x):
            count[x]-=1
            if count[x]==0: count.pop(x)
        def updatecount():
            nonlocal count
            count=Counter()
            if len(sl)==0: return
            elif len(sl)==1: 
                count[n]=1
                return
            for ind in range(len(sl)):
                count[getsl(ind)-getsl(ind-1)]+=1
        updatecount()
        def toggle(i):
            if i<0: i+=n
            if i>=n: i-=n
            if len(sl)<5:
                if i in sl: sl.remove(i)
                else: sl.add(i)
                updatecount()
                return
            if i in sl:
                ind=sl.bisect_left(i)
                previ,nexti = getsl(ind-1), getsl(ind+1)
                dec(i-previ)
                dec(nexti-i)
                inc(nexti-previ)
                sl.remove(i)
            else:
                ind=sl.bisect_left(i)
                previ,nexti=getsl(ind-1), getsl(ind)
                inc(i-previ)
                inc(nexti-i)
                dec(nexti-previ)
                sl.add(i)
        res=[]
        for q in queries:
            if q[0]==1:
                size=q[1]
                if len(sl)==0: res.append(n)
                else: res.append(sum(v*max(0, k-size+1) for k,v in count.items()))
            else:
                i,col = q[1], q[2]
                if colors[i]==col: continue
                colors[i]=col
                toggle(i-1)
                toggle(i)
        return res
    def longestCommonSubpath(self, n: int, paths: List[List[int]]) -> int:
        arr=[]
        m_path_index=[]
        for i,p in enumerate(paths):
            arr.extend(p)
            arr.append(-i-1)
            m_path_index.extend([i]*(len(p)+1))
        suffix_array=build_suffix_array(arr)
        lcp_array=build_lcp_array(arr, suffix_array)
        res,c,dq,start = 0,Counter(),deque(),0
        for i,v in enumerate(lcp_array):
            c[m_path_index[suffix_array[i]]] += 1
            while dq and lcp_array[dq[-1]] >=v: dq.pop()
            dq.append(i)
            while len(c)==len(paths):
                if start==dq[0]: dq.popleft()
                res=max(res, lcp_array[dq[0]])
                x = m_path_index[suffix_array[start]]
                c[x]-=1
                if c[x]==0: c.pop(x)
                start+=1
        return res



            
class Manacher:
    def __init__(self, s):
        a='#' + '#'.join(s)+'#'
        n=len(a)
        data=[0]*len(a)
        lo,hi=0,0
        for i in range(n):
            if hi>i: data[i] = min(hi-i, data[lo+hi-i])
            while i-data[i]>=0 and i+data[i]<n and a[i-data[i]]==a[i+data[i]]: data[i]+=1
            if i+data[i]>hi: lo,hi = i-data[i], i+data[i]
        print(a, data)
        def query(i,j):
            return data[i+j+1]>j-i+1
        self.query=query


def build_lcp_array(s, suffix_array):
    n=len(s)
    suffix_index=[0]*n
    for i,v in enumerate(suffix_array): suffix_index[v]=i
    lcp=[0]*n
    k=0
    for i in range(n):
        ind = suffix_index[i]
        if ind>0:
            j=suffix_array[ind-1]
            while i+k<n and j+k<n and s[i+k]==s[j+k]: k+=1
            lcp[ind] = k
        if k>0: k-=1
    return lcp


class LCPQuery:
    def __init__(self, s):
        suffix_array = build_suffix_array(s)
        n=len(s)
        suffix_index=[0]*n
        for i,v in enumerate(suffix_array): suffix_index[v]=i
        lcp=[0]*n
        k=0
        for i in range(n):
            ind = suffix_index[i]
            if ind>0:
                j=suffix_array[ind-1]
                while i+k<n and j+k<n and s[i+k]==s[j+k]: k+=1
                lcp[ind] = k
            if k>0: k-=1
        rmq = RMQ(lcp)
        def query(i,j):
            if i==j: return n-i
            ii,jj = suffix_index[i], suffix_index[j]
            if ii>jj: ii,jj=jj,ii
            return lcp[rmq.query(ii+1, jj)]
        self.query = query


def countpre(a,b):
    n=min(len(a), len(b))
    for i in range(n):
        if a[i]!=b[i]: return i
    return n

s = '3712406596575692459375265479356465193982807374167'
lcpq = LCPQuery(s)
print("checking", s)
for i,j in combinations(range(len(s)), 2):
    if lcpq.query(i,j) != countpre(s[i:], s[j:]):
        print(i,j, s[i:], s[j:], lcpq.query(i,j))


class FenwickTree2D:
    def __init__(self, n):
        self.n = n
        self.BIT = [[0] * (n + 1) for _ in range(n + 1)]
        self.grid = [[0] * n for _ in range(n)]

    def _update(self, i, j, delta):
        x = i + 1
        while x <= self.n:
            y = j + 1
            while y <= self.n:
                self.BIT[x][y] += delta
                y += y & -y
            x += x & -x
    
    def update(self, i, j, v):
        delta = v - self.grid[i][j]
        self.grid[i][j] = v
        self._update(i, j, delta)

    def _sum(self, i, j):
        total = 0
        x = i + 1
        while x > 0:
            y = j + 1
            while y > 0:
                total += self.BIT[x][y]
                y -= y & -y
            x -= x & -x
        return total

    def sumrange(self, i, j):
        return self._sum(i, j)

MOD = 10**9 + 7
@lru_cache(None)
def fac(v):
    if v==0: return 1
    return v * fac(v-1)%MOD
def mod_inverse(v):
    return pow(v, MOD-2, MOD)
def comb(n,k):
    return fac(n) * mod_inverse(fac(k) * fac(n-k)) % MOD