package main

func permute(nums []int) [][]int {
	var result [][]int
	var current []int
	if len(nums) <= 1 {
		result = append(result, nums)
		return result
	}
	for i := range nums {
		current_copy := make([]int, len(current))
		copy(current_copy, current)
		current_copy = append(current_copy, nums[i])
		tmp := make([]int, i)
		copy(tmp, nums[0:i])
		subPermute(append(tmp, nums[i+1:len(nums)]...), current_copy, &result)
	}
	return result
}

func subPermute(nums []int, current []int, result *[][]int) {
	if len(nums) == 1 {
		current = append(current, nums[0])
		tmp := make([]int, len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
		return
	}
	for i := range nums {
		current_copy := make([]int, len(current))
		copy(current_copy, current)
		current_copy = append(current_copy, nums[i])
		tmp := make([]int, i)
		copy(tmp, nums[0:i])
		subPermute(append(tmp, nums[i+1:len(nums)]...), current_copy, result)
	}
	return
}

func main() {}
