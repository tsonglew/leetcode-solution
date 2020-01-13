func missingNumber(nums []int) int {
    sm := 0
	for _, n := range nums {
		sm += n
	}
	return len(nums)*(len(nums)+1)/2 - sm
}
