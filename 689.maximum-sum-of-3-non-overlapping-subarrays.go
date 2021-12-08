func maxSumOfTwoSubarrays(nums []int, k int) (ans []int) {
    var sum1, maxSum1, maxSum1Idx int
    var sum2, maxSum12 int
    for i := k; i < len(nums); i++ {
        sum1 += nums[i-k]
        sum2 += nums[i]
        if i >= k*2-1 {
            if sum1 > maxSum1 {
                maxSum1 = sum1
                maxSum1Idx = i - k*2 + 1
            }
            if maxSum1+sum2 > maxSum12 {
                maxSum12 = maxSum1 + sum2
                ans = []int{maxSum1Idx, i - k + 1}
            }
            sum1 -= nums[i-k*2+1]
            sum2 -= nums[i-k+1]
        }
    }
    return
}
