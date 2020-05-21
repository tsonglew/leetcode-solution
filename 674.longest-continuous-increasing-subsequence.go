func findLengthOfLCIS(nums []int) int {
    if len(nums) < 2 { return len(nums) }
    ans := 1
    cur := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            cur++
            if cur > ans {
                ans = cur
            }
        } else {
            cur = 1
        }
    }
    return ans
}
