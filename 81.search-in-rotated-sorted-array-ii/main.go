package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 1, 3, 1}, 3))
	fmt.Println(search([]int{2, 5, 6, 1, 2}, 5))
}

func search(nums []int, target int) bool {
	for lo, hi := 0, len(nums)-1; lo <= hi; {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < target && target < nums[hi] {
			lo = mid + 1
		} else if target < nums[mid] && nums[lo] < target {
			hi = mid - 1
		} else {
			if nums[lo] == target || nums[hi] == target {
				return true
			} else {
				lo++
				hi--
			}
		}
	}
	return false
}
