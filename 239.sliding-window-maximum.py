import heapq

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        result = []
        if len(nums) == 0 or k == 0:
            return result
        i = 0
        while i <= len(nums) - k:
            result.append(heapq.nlargest(1, nums[i:i+k])[0])
            i += 1
        return result