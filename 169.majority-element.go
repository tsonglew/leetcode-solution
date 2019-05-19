// hashtable
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		v, ok := m[nums[i]]
		if ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
		if v == (len(nums)-1)/2 {
			return nums[i]
		}
	}
	return 0
}

// moore voting
func majorityElement(nums []int) int {
	cnt := 0
	major := 0
	for i := range nums {
		if cnt == 0 {
			major = nums[i]
			cnt++
		} else if nums[i] == major {
			cnt++
		} else if nums[i] != major {
			cnt--
		}
	}
	return major
}