package main

func main() {
	print(rob([]int{1, 2, 1, 0}))
}

func rob(nums []int) int {
	if len(nums) <= 3 {
		return max(nums...)
	}
	return max(rob2(append([]int{}, nums[1:]...)), rob2(append([]int{}, nums[:len(nums)-1]...)))
}

func rob2(nums []int) int {
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}
	return nums[len(nums)-1]
}

func max(nums ...int) int {
	mx := 0
	for _, n := range nums {
		if n > mx {
			mx = n
		}
	}
	return mx
}
