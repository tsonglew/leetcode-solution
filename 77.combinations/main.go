package main

func main() {

}

func combine(n int, k int) [][]int {
	result := [][]int{}
	comb := []int{}
	subcom(0, k, n, &comb, &result)
	return result
}

func subcom(s, k, n int, comb *[]int, result *[][]int) {
	if k > 0 {
		for i := s + 1; i <= n-k+1; i++ {
			c := append(*comb, i)
			subcom(i, k-1, n, &c, result)
		}
	} else {
		*result = append(*result, append([]int{}, *comb...))
	}
}
