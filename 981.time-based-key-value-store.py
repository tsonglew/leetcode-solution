from collections import defaultdict

class TimeMap:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.map = defaultdict(list)


    def set(self, key: str, value: str, timestamp: int) -> None:
        self.map[key].append((value, timestamp))

    @staticmethod
    def binsearch(l, timestamp):

        low, high = 0, len(l) -1
        while low < high:
            mid = (low+high)//2
            if l[mid][1] > timestamp:
                high = mid-1
            elif l[mid][1] < timestamp:
                low = mid+1
            else:
                return l[mid][0]
        while low >= 0 and low < len(l):
            if l[low][1] <= timestamp:
                return l[low][0]
            low -= 1
        return ""



    def get(self, key: str, timestamp: int) -> str:
        return self.binsearch(self.map[key], timestamp)



# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)
