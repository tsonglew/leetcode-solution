import heapq
from collections import Counter

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        nums_count = dict()
        for i in nums:
            if i in nums_count.keys():
                nums_count[i][1] += 1
            else:
                nums_count[i] = [i, 1]
        return [x[0] for x in heapq.nlargest(k, nums_count.values(), key=lambda x: x[1])]

    def topKFrequentCounter(self, nums: List[int], k: int) -> List[int]:
        return [num[0] for num in Counter(nums).most_common(k)]