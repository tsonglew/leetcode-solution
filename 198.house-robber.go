func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[1], nums[0])
	case 3:
		return max(nums[1], nums[0]+nums[2])
	}
	nums[2] = nums[0] + nums[2]
	for i := 3; i < len(nums); i++ {
		nums[i] = max(nums[i]+nums[i-2], nums[i]+nums[i-3])
	}
	return max(nums[len(nums)-3:]...)
}

func max(nums ...int) int {
	mx := nums[0]
	for i := range nums {
		if nums[i] > mx {
			mx = nums[i]
		}
	}
	return mx
}