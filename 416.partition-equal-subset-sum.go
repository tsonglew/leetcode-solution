func canPartition(nums []int) bool {
    sum := sumInts(nums)
    half, remain := sum / 2, sum % 2
    if remain != 0 {
        return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    if canSumTo(nums, half, 0, half, true) {
        return true
    }
    return false
}

func sumInts(nums []int) int {
    sum := 0
    for i := range nums {
        sum += nums[i]
    }
    return sum
}

func canSumTo(nums []int, sum, minused, half int, first bool) bool {
    if sum == 0 {
        return true
    }
    if sum < 0 {
        return false
    }
	tried := make(map[int]bool)
    for i := range nums {
		if first {
			if b, ok := tried[nums[i]]; ok && !b {
				continue
			}
		}
		if minused > half {
			return false
		}
        if canSumTo(append([]int{}, nums[i+1:len(nums)]...), sum - nums[i], minused, half, false) {
            return true
		}
		minused += nums[i]
        if first {
            tried[nums[i]] = false
        }
    }
    return false
}