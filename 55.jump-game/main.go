package main

import "fmt"

func main(){
	fmt.Println(spiralOrder([][]int{}))
}

func canJump(nums []int) bool {
	m := make([]bool, len(nums))
	m[0] = true
	for i := range nums {
		if !m[i] {
			continue
		}
		for j := nums[i];j>0;j-- {
			if i+j >= len(nums) {
				j = len(nums) - i - 1
			}
			m[i+j] = true
		}
		if m[len(nums)-1] == true {
			return true
		}
	}
	return false
}