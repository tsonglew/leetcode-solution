class Solution:
    def reversePairs(self, nums: List[int]) -> int:
        return self.mergeSort(nums)[0]

    def mergeSort(self, nums):
        if len(nums) <= 1:
            return 0, nums

        mid = len(nums) // 2
        res_left, sorted_left = self.mergeSort(nums[:mid])
        res_right, sorted_right = self.mergeSort(nums[mid:])
        res = res_left + res_right

        sorted_nums = []
        i, j = 0, 0

        while i < len(sorted_left) and j < len(sorted_right):
            if sorted_left[i] <= sorted_right[j]:
                sorted_nums.append(sorted_left[i])
                i += 1
            else:
                sorted_nums.append(sorted_right[j])
                res += len(sorted_left) - i
                j += 1
        if i < len(sorted_left):
            sorted_nums.extend(sorted_left[i:])
        if j < len(sorted_right):
            sorted_nums.extend(sorted_right[j:])
        return res, sorted_nums
