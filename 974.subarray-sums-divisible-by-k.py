class Solution:
    def subarraysDivByK(self, A: List[int], K: int) -> int:
        from collections import defaultdict
        ans = 0
        d = defaultdict(lambda: 0)
        d[0] = 1
        total = 0
        for i in A:
            total = (total+i)%K
            ans += d[total]
            d[total] += 1
        return ans
