func findMin(nums []int) int {
    if nums[0] < nums[len(nums)-1] {
        return nums[0]
    }
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
    if nums[mid] < nums[lo] {
        return find(nums, lo, mid)
    }
    return find(nums, mid+1, hi)
}