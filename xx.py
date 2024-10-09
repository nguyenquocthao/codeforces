from sortedcontainers import SortedList
from typing import List
from collections import defaultdict
def run(a:List[int], b:List[int], queries):
    m, l = defaultdict(SortedList), []
    for i,v in enumerate(b):
        