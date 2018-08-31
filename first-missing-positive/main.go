package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	if len(nums) < 1 {
		return 1
	}
	tmp := make([]int, len(nums)+1)
	for _, v := range nums {
		if v > 0 && v <= len(nums) {
			tmp[v-1] = 1
		}
	}
	for i := range tmp {
		if tmp[i] == 0 {
			return i + 1
		}
	}
	return 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1}))
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
