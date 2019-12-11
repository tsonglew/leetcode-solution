func rotate(nums []int, k int)  {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, s, e int) {
	for i, j := s, e; i < j; i, j = i+1, j-1{
		nums[i], nums[j] = nums[j], nums[i]
	}
}