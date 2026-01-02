from itertools import permutations, combinations, product
from collections import defaultdict,Counter
import random
from math import gcd, lcm
from functools import lru_cache, cmp_to_key
from typing import List

import sys


# Check current recursion limit
print(sys.getrecursionlimit())

# Increase recursion limit, e.g. to 5000
sys.setrecursionlimit(50000)

# Confirm change
print(sys.getrecursionlimit())


MAXV = 8
REPEAT=3
@lru_cache(None)
def nb(tup):
    l = list(tup)
    res = set()
    for v in range(MAXV):
        for i,j in combinations(range(len(l)), 2):
            x,y = l[i], l[j]
            l[i], l[j] = x&v, y|v
            res.add(tuple(l))
            l[i], l[j] = x|v, y&v
            res.add(tuple(l))
            l[i], l[j] = x,y
    return sorted(res)

def build_graph():
    """Build adjacency list for all tuples."""
    graph = defaultdict(list)
    for tup in product(range(MAXV), repeat=REPEAT):
        graph[tup] = nb(tup)
    return graph

def tarjan_scc(graph):
    """Tarjan’s algorithm to find strongly connected components."""
    index = 0
    stack = []
    indices = {}
    lowlink = {}
    onstack = set()
    sccs = []

    def strongconnect(node):
        nonlocal index
        indices[node] = index
        lowlink[node] = index
        index += 1
        stack.append(node)
        onstack.add(node)

        for neighbor in graph[node]:
            if neighbor not in indices:
                strongconnect(neighbor)
                lowlink[node] = min(lowlink[node], lowlink[neighbor])
            elif neighbor in onstack:
                lowlink[node] = min(lowlink[node], indices[neighbor])

        # root node of SCC
        if lowlink[node] == indices[node]:
            comp = []
            while True:
                w = stack.pop()
                onstack.remove(w)
                comp.append(w)
                if w == node:
                    break
            sccs.append(comp)

    for node in graph:
        if node not in indices:
            strongconnect(node)
    return sccs

def build_dag(graph, sccs):
    """Build DAG where each SCC is a single node."""
    # map each node to its SCC index
    node_to_scc = {}
    for i, comp in enumerate(sccs):
        for node in comp:
            node_to_scc[node] = i

    dag = defaultdict(set)
    for u in graph:
        for v in graph[u]:
            if node_to_scc[u] != node_to_scc[v]:
                dag[node_to_scc[u]].add(node_to_scc[v])
    return dag

# if __name__ == "__main__":
#     graph = build_graph()
#     sccs = tarjan_scc(graph)
#     dag = build_dag(graph, sccs)

#     print(f"Total nodes: {len(graph)}, SCC count: {len(sccs)}")
#     for i, comp in enumerate(sccs):
#         print(f"SCC {i}: {[tuple(bin(x)[2:].zfill(4) for x in t) for t in comp]}")


#     print("\nDAG edges:")
#     for u in dag:
#         print(f"{u} -> {sorted(dag[u])}")
    
            


# # c=0
# k=2
# la, lb = [], []
# for i in range(1, 10**7):
#     x,y = str(i), cvk(i,k)
#     if x==x[::-1]: la.append(x)
#     if y==y[::-1]: lb.append(y)
#     # if iskmirror(i, k):
#     #     x, y = str(i), cvk(i, k)
#     #     c+=1
#     #     print(dict(c=c,x=x,y=y))

# print(len(la), len(lb))
# print(la[:10])
# print(lb[:10])


MOD = 10**9 + 7
@lru_cache(None)
def p2(n):
    return pow(2, n, MOD)
class Solution:
    def countEffective(self, nums: List[int]) -> int:
        orv=0
        for v in nums: orv|=v
        if orv==0: return 0
        nb = orv.bit_length()
        count = [0]*(orv+1)
        for v in nums: count[v]+=1
        # # 111 => 111, 011, 101, 110, 100, 010, 001, 000
        for b in range(nb):
            mask=(1<<b)
            if mask&orv==0: continue
            for i,v in enumerate(count):
                if v>0 and i&mask==0: count[i+mask] += v
        res=0
        for i,v in enumerate(count):
            if i|orv!=orv: continue
            # if not v: continue
            nb = i.bit_count()
            res+=(-1)**nb * p2(v)
        # print(count, orv, res)
        # print(count)
        return (p2(len(nums)) - abs(res))%MOD
    
    

ll=[]
for i,j in product(range(22),repeat=2):
    k = i^j
    while k>=len(ll): ll.append([])
    ll[k].append((i,j))
for k,l in enumerate(ll):
    print(k,len(l),l)


        