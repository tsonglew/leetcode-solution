func findPeakElement(nums []int) int {
    return find(nums, 0, len(nums)-1)
}

func find(nums []int, lo, hi int) int {
    if lo == hi {
        return lo
    }
    mid := (lo + hi) / 2
    if nums[mid] > nums[mid+1] {
        return find(nums, lo, mid)
    }
    return find(nums, mid+1, hi)
}