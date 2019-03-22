package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{[]int{}}
	f(0, nums, []int{}, &result)
	return result
}

func f(n int, nums, current []int, result *[][]int) {
	for i := n; i < len(nums); i++ {
		if i > n && nums[i] == nums[i-1] {
			continue
		}
		appendedCurrent := append([]int{}, current...)
		appendedCurrent = append(appendedCurrent, nums[i])
		*result = append(*result, appendedCurrent)
		f(i+1, nums, appendedCurrent, result)
	}
}
