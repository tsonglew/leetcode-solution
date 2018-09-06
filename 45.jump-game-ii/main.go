package main

import (
	"math"
)

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	jumps := make([]int, len(nums))
	jumps[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			jumps[i] = math.MaxInt64
			continue
		}
		maxIndex := nums[i] + i
		if maxIndex > len(nums)-1 {
			maxIndex = len(nums) - 1
		}
		minJumps := jumps[maxIndex]
		for j := maxIndex; j > i; j-- {
			if jumps[j] < minJumps {
				minJumps = jumps[j]
			}
		}
		if minJumps == math.MaxInt64 {
			jumps[i] = math.MaxInt64
		} else {
			jumps[i] = minJumps + 1
		}
	}
	return jumps[0]
}

func main() {

}
