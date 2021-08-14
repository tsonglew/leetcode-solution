class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        res = 0
        cnt = 0
        for i in nums:
            if res == i:
                cnt += 1
            else:
                if cnt > 0:
                    cnt -= 1
                else:
                    res = i
                    cnt = 1

        return res
