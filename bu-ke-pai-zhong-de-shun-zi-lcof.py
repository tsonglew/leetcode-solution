class Solution:
    def isStraight(self, nums: List[int]) -> bool:
        zeros = 0
        nums.sort()
        for i in range(1, len(nums)):
            if nums[i-1] == 0:
                zeros += 1
                continue
            if nums[i] == nums[i-1]:
                return False
            if nums[i] != nums[i-1]+1:
                gap = nums[i] - nums[i-1] - 1
                if zeros >= gap:
                    zeros -= gap
                    continue
                return False
        return True
