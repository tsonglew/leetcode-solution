func productExceptSelf(nums []int) []int {
    if len(nums) == 0 {
        return nums
    }
    if len(nums) == 1 {
        return []int{0}
    }
    left := make([]int, len(nums))
    right := make([]int, len(nums))
    left[0] = 1
    right[len(nums)-1] = 1
    for i := 1; i < len(nums); i++ {
        left[i] = left[i-1]*nums[i-1]
        right[len(nums)-1-i] = right[len(nums)-i]*nums[len(nums)-i]
    }
    for i := range left {
        left[i] *= right[i]
    }
    return left
}

