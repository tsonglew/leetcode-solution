class Solution:
    def isEven(self, num: int) -> int:
        cnt = 0
        while num > 0:
            num //= 10
            cnt += 1
        return cnt % 2 == 0

    def findNumbers(self, nums: List[int]) -> int:
        cnt = 0
        for n in nums:
            if self.isEven(n):
                cnt += 1
        return cnt
