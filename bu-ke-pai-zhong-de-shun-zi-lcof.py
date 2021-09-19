class Solution:
    def isStraight(self, nums: List[int]) -> bool:
        nums.sort()
        zeros = len(list(filter(lambda x: x==0, nums)))
        for i in range(zeros+1, len(nums)):
            if nums[i] == nums[i-1]:
                return False
            if nums[i] != nums[i-1]+1:
                gap = nums[i] - nums[i-1] - 1
                if zeros >= gap:
                    zeros -= gap
                    continue
                return False
        return True
