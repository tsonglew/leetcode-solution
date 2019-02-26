package main

func main() {

}

func subsets(nums []int) [][]int {
	result := [][]int{}
	for i := 0; i <= len(nums); i++ {
		result = append(result, combine(nums, i)...)
	}
	return result
}

func combine(nums []int, k int) [][]int {
	result := [][]int{}
	comb := []int{}
	subcom(0, k, nums, &comb, &result)
	return result
}

func subcom(s, k int, nums []int, comb *[]int, result *[][]int) {
	if k > 0 {
		for i := s; i < len(nums)-k+1; i++ {
			c := append(*comb, nums[i])
			subcom(i+1, k-1, nums, &c, result)
		}
	} else {
		*result = append(*result, append([]int{}, *comb...))
	}
}
