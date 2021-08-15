class Solution:
    def getLeastNumbers(self, arr: List[int], k: int) -> List[int]:
        return heapq.nsmallest(k, arr)
