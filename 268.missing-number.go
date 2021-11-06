func missingNumber(nums []int) int {
    r := (len(nums)+1)*len(nums)/2
    for i := range nums {
        r -= nums[i]
    }
    return r
}
