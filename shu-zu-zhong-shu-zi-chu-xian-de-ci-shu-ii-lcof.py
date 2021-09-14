class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        ones = 0
        twos = 0
        for n in nums:
            twos, ones = (twos&~n)|(~twos&n&ones), ~twos&(ones^n)
        return ones
