func numTrees(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	case 2:
		return 2
	default:
		return catalan(n)
	}
}

func catalan(n int) int {
	nums := make([]int, n+1)
	nums[0], nums[1], nums[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		sum := 0
		for j := 0; j < i/2; j++ {
			sum += 2 * nums[j] * nums[i-j-1]
		}
		if i%2 == 1 {
			sum += nums[i/2] * nums[i/2]
		}
		nums[i] = sum
	}
	return nums[n]
}