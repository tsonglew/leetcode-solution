// heap like
func findKthLargest(nums []int, k int) int {
	for c := k; c > 0; c-- {

		for i := len(nums) - 1; i >= 0; i-- {
			sortChild(nums, i)
		}
		if c == 1 {
			return nums[0]
		}
		nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
		nums = nums[0 : len(nums)-1]
	}
	return -1
}

func sortChild(nums []int, k int) {
	leftChildIdx := k*2 + 1
	rightChildIdx := k*2 + 2
	nextIdx := k
	if leftChildIdx < len(nums) && nums[nextIdx] < nums[leftChildIdx] {
		nextIdx = leftChildIdx
	}
	if rightChildIdx < len(nums) && nums[nextIdx] < nums[rightChildIdx] {
		nextIdx = rightChildIdx
	}
	if nextIdx == k {
		return
	}
	nums[nextIdx], nums[k] = nums[k], nums[nextIdx]
	sortChild(nums, nextIdx)
}

// quicksort like
