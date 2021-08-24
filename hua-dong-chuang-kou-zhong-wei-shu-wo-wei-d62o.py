class MedianFinder:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.lower = []
        self.higher = []


    def addNum(self, num: int) -> None:
        if len(self.lower) == len(self.higher):
            heapq.heappush(self.higher, -heapq.heappushpop(self.lower, -num))
        else:
            heapq.heappush(self.lower, -heapq.heappushpop(self.higher, num))


    def findMedian(self) -> float:
        if len(self.lower) == len(self.higher):
            return (-self.lower[0] + self.higher[0]) / 2
        return self.higher[0]



# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
