func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    for i := range nums {
		for j := i+1; j<=i+k&&j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}