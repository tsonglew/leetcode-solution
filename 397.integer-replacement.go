func integerReplacement(n int) int {
	return subReplace(n, 0)
}

func subReplace(n, t int) int {
	if n == 1 {
		return t
	}
	if n % 2 == 1 {
		return min(subReplace(n-1, t+1), subReplace(n+1, t+1))
	}
	return subReplace(n/2, t+1)
}

func min(nums ...int) int {
	mi := nums[0]
	for i := range nums {
		if nums[i] < mi {
			mi = nums[i]
		}
	}
	return mi
}