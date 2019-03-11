package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j, repeatCnt, currentNum := 0, 0, 2, nums[0]
	for i < len(nums) && j < len(nums) {
		if nums[j] == currentNum {
			if repeatCnt != 0 {
				nums[i] = currentNum
				repeatCnt--
				i++
			}
			j++
		} else {
			currentNum = nums[j]
			repeatCnt = 2
		}
	}
	return i
}
