class NumArray:

    def __init__(self, nums: List[int]):
        self.nums_sum = [0]*len(nums)
        for i, _ in enumerate(nums):
            self.nums_sum[i] = self.nums_sum[i-1] + (nums[i] if i > 0 else nums[i])

    def sumRange(self, left: int, right: int) -> int:
        return self.nums_sum[right] - self.nums_sum[left] + self.nums_sum[left] - (self.nums_sum[left-1] if left > 0 else 0)
