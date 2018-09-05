package main

func searchRange(nums []int, target int) []int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) []int {
	if end == -1 || (start == end && nums[start] != target) {
		return []int{-1, -1}
	}
	half := (end + start) / 2
	if nums[half] < target {
		start = half + 1
	} else if nums[half] > target {
		end = half - 1
	} else {
		ft := firstTarget(nums, start, half, target)
		lt := lastTarget(nums, half, end, target)
		return []int{ft, lt}
	}
	return binarySearch(nums, start, end, target)
}

func firstTarget(nums []int, start, end, target int) int {
	half := (start + end) / 2
	if nums[half] == target {
		if half == start {
			return start
		}
		return firstTarget(nums, start, half, target)
	}
	return firstTarget(nums, half+1, end, target)
}

func lastTarget(nums []int, start, end, target int) int {
	half := (start + end + 1) / 2
	if nums[half] == target {
		if half == start {
			return start
		}
		return lastTarget(nums, half, end, target)
	}
	return lastTarget(nums, half, end-1, target)
}

func main() {}
