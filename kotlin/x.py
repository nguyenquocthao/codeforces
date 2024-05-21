from collections import defaultdict
from typing import List

class Tree:
    def __init__(self, parent:List[int], mask:int):
        n = len(parent)
        mp, mc, mtrue = {}, defaultdict(list), list(range(n)) 
        def gettrue(i):
            if i==0: return 0
            if mask&(1<<mtrue[i]) == 0: 
                mtrue[i] = gettrue(parent[i])
            return mtrue[i]
        for i,p in enumerate(parent):
            if i==0 or mask&(1<<i) == 0: continue
            p = gettrue(p)
            mp[i] = p
            mc[p].append(i)
        self.mparent, self.mchild = mp, mc
    def get_descendants(self, i:int) -> int:
        res = 1<<i
        for j in self.mchild[i]: res |= self.get_descendants(j)
        return res
    def get_ancestors(self, i:int) -> int:
        res = 0
        while True:
            i = self.mparent[i]
            if i==0: break
            res |= 1<<i            
        return res
    
def run(a:List[int], b:List[int]):
    a,b = [0] + [v-1 for v in a], [0] + [v-1 for v in b]
    n = len(a)
    def dp(mask):
        if mask.bit_count()<=2: return 0
        treea, treeb = Tree(a, mask), Tree(b,mask)
        # if len(treea.mchild[0])==0: return 0
        x = treea.mchild[0][0]
        mask -= 1<<x
        desa, desb = treea.get_descendants(x) - (1<<x), treeb.get_descendants(x) - (1<<x)
        rm = treeb.get_ancestors(x) | (desa ^ desb)
        if rm == 0:
            # only descentdant of x / other
            return dp(desa) + dp(mask - desa)
        res = rm.bit_count()
        common = desa & desb
        res += dp(common) + dp(mask & ~rm & ~common)
        return min(res, 1 + dp(mask))
    return dp((1<<n)-1)

# print(2*run([4, 7, 10, 6, 2, 9, 7, 1, 1], [1, 2, 10, 3, 4, 6, 6, 5, 5]))

with open('x.txt','w') as f:
    n, q = 200000, 200000
    print(n,q, file=f)
    print(' '.join(['10000']*n), file=f)
    print(' '.join(['1']*n), file=f)
    for i in range(q):
        print(i+1, n, file=f)









     

                
