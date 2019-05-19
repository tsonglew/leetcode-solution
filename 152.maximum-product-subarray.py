class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        r = nums[0]
        cmax, cmin = r, r
        for i in nums[1:]:
            if i < 0:
                cmax, cmin = cmin, cmax
            cmax = max(cmax * i, i)
            cmin = min(cmin * i, i)
            r = max(cmax, r)
        return r