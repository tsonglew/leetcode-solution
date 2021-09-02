class Solution:
    def missingNumber(self, nums) -> int:
        if nums[0] != 0:
            return 0
        if nums[-1] != len(nums):
            return len(nums)
        return self.f(nums)
        
    def f(self, nums):
        print(nums)
        if len(nums) <= 3:
            for i in range(1, len(nums)):
                if nums[i] != nums[i-1] + 1:
                    return nums[i-1] + 1
        mid = len(nums) // 2
        if nums[mid] != nums[0] + mid:
            return self.f(nums[:mid+1])
        return self.f(nums[mid:])
