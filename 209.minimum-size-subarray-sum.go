func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minLen := len(nums)+1
	sum := nums[0]
	for i, j := 0, 1; j < len(nums); {
		for ;sum < s && j < len(nums); j++ {
			sum += nums[j]
		}
		for ;sum >= s; i++ {
			if j-i < minLen {
				minLen = j-i
			}
			sum -= nums[i]
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}