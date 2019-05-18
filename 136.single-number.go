func singleNumber(nums []int) int {
	sum := 0
	for i := range nums {
		sum ^= nums[i]
	}
	return sum
}