func subarraySum(nums []int, k int) int {
    total := 0
    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i-1]
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] == k {
            total++
        }
        for j := 0; j < i; j++ {
            if nums[i] - nums[j] == k {
                total++
            }
        }
    }
    return total
}