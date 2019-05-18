type scope struct {
	low, high int
}

func longestConsecutive(nums []int) int {
	m := make(map[int]scope)
	maxCons := 0
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		lower, higher := nums[i], nums[i]
		for s, ok := m[nums[i]-1]; ok; {
			lower = s.low
			s, ok = m[lower-1]
		}
		for s, ok := m[nums[i]+1]; ok; {
			higher = s.high
			s, ok = m[higher+1]
		}
		cons := higher - lower + 1
		if cons > maxCons {
			maxCons = cons
		}
		m[nums[i]] = scope{lower, higher}
	}
	return maxCons
}