class Solution:
    def isEven(self, num: int) -> int:
        cnt = 0
        while num > 0:
            num //= 10
            cnt += 1
        return cnt % 2 == 0

    def findNumbersNaive(self, nums: List[int]) -> int:
        cnt = 0
        for n in nums:
            if self.isEven(n):
                cnt += 1
        return cnt
    
    def findNumbersWithMath(self, nums: List[int]) -> int:
        from math import log10
        cnt = 0
        for n in nums:
            if int(log10(n)+1) % 2 == 0:
                cnt += 1
        return cnt
