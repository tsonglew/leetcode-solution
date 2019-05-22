import heapq

class Solution(object):

    # use heap
    def findKthLargestH(self, nums, k):
        nums = [-num for num in nums]
        heapq.heapify(nums)
        res = float('inf')
        for _ in range(k):
            res = heapq.heappop(nums)
        return -res

    # use sort
    def findKthLargestH(self, nums, k):
        return sorted(nums)[-k]
