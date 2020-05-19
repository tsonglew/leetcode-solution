func search(nums []int, target int) int {
    return subSearch(nums, 0, len(nums)-1, target)
}

func subSearch(nums []int, lo, hi int, target int) int {
    if lo > hi {
        return -1
    }
    mid := (lo+hi)/2
    if nums[mid] == target {
        return mid
    }
    if nums[lo] <= nums[mid] {
        if nums[lo] <= target && target <= nums[mid] {
            return subSearch(nums, lo, mid-1, target)
        }
        return subSearch(nums, mid+1, hi, target)
    }
    if nums[mid] <= target && target <= nums[hi] {
        return subSearch(nums, mid+1, hi, target)
    }
    return subSearch(nums, lo, mid-1, target)
}
