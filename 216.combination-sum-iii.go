func combinationSum3(k int, n int) [][]int {
	return combine(1, k, n)
}

func combine(s, k, n int) [][]int {
	if k == 0 {
		if n == 0 {
			return [][]int{[]int{}}
		}
		return [][]int{}
	}
	result := [][]int{}
	for i := s; i <= 10-k && n >= i; i++ {
		combined := combine(i+1, k-1, n-i)
		for j := range combined {
			result = append(result, append([]int{i}, combined[j]...))
		}
	}
	return result
}