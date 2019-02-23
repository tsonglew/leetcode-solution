package main

func main() {

}

func searchLeast(nums []int, position, least int) (int, int) {
	numsLen := len(nums)
	min := nums[position]
	minPosition := position
	for i := position; i < numsLen; i++ {
		if nums[i] < min {
			min = nums[i]
			minPosition = i
		}
		if nums[i] == least {
			return least, i
		}
	}
	return min, minPosition
}

func sortColors(nums []int) {
	least := 0
	j := 0
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if nums[i] == least {
			continue
		} else {
			if i < numsLen-1 {
				least, j = searchLeast(nums, i, least)
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		if least == 2 {
			return
		}
	}
}
