func findMin(nums []int) int {
    return find(nums, 0, len(nums)-1)
}

func find(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}
	mid := (hi + lo) / 2
	if nums[mid] > nums[mid+1] {
		return nums[mid+1]
	}
	return min(find(nums, lo, mid), find(nums, mid+1, hi))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}