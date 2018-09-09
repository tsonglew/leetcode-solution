package maxsubarray

func maxSubArray(nums []int) int {
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		sum := nums[i-1] + nums[i]
		if sum > nums[i] {
			nums[i] = sum
		}
		if nums[i] > mx {
			mx = nums[i]
		}
	}
	return mx
}
