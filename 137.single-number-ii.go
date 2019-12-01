package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber([]int{1}))
}

func singleNumber(nums []int) int {
	// https://blog.csdn.net/qq_41231926/article/details/85758627
	ones, twos, mask := 0, 0, 0
	for i := range nums {
		twos ^= nums[i] & ones
		ones ^= nums[i]
		mask = ^(twos & ones)
		ones &= mask
		twos &= mask
	}
	return ones
}
