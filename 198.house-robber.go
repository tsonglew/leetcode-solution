func rob(nums []int) int {
    ans := 0
    for i := range nums {
        cur := nums[i]
        if i >= 2 && nums[i-2] + nums[i] > cur {
            cur = nums[i-2] + nums[i]
        }
        if i >= 3 && nums[i-3] + nums[i] > cur {
            cur = nums[i-3] + nums[i]
        }
        nums[i] = cur
        if nums[i] > ans {
            ans = nums[i]
        }
    }
    return ans
}
