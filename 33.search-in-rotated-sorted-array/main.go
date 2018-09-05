package main

func search(nums []int, target int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}
	lo := 0
	hi := length - 1
	var mid int
	for lo < hi {
		mid = (lo + hi) / 2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	hi = (lo - 1 + length) % length
	lf := binarySearch(nums, 0, hi, target)
	rt := binarySearch(nums, lo, length-1, target)
	if lf != -1 {
		return lf
	}
	return rt
}

func binarySearch(nums []int, start, end, target int) int {
	mid := (start + end) / 2
	for start < end {
		mid = (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			break
		}
	}
	if start == end {
		if nums[start] != target {
			return -1
		}
		return start
	}
	if nums[mid] == target {
		return mid
	}
	return -1
}

func main() {}
