isprime = [True]*1001
isprime[0], isprime[1] = False, False
for i in range(2, len(isprime)):
    if isprime[i]:
        for j in range(i*i, len(isprime), i):
            isprime[j]=False
lp = [i for i,ok in enumerate(isprime) if ok]
acc = list(accumulate(lp))

@lru_cache(None)
def dp(i,s):
    if i<0: return []
    if s>=acc[i]: return lp[:i+1]
    v = lp[i]
    if v>s: return dp(i-1, s)
    if v==s: return [v]
    x,y = dp(i-1, s), dp(i-1, s-v) + [v]
    if (-sum(x), len(x), x) < (-sum(y), len(y), y):
        return x
    else:
        return y
class Solution:
    def maxPrimeSubset(self, n: int, s: int) -> list[int]:
        return dp(bisect.bisect_right(lp, n) - 1, s)
        