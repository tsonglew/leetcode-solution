func summaryRanges(nums []int) []string {
	result := []string{}
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		last := nums[i]
		var j int
		for j = i; j <= len(nums)-1; j++ {
			if j >= len(nums) - 1 || nums[j+1] != last+1 {
				if start != last {
					result = append(result, fmt.Sprintf("%d->%d", start, last))
				} else {
					result = append(result, fmt.Sprintf("%d", start))
				}
				i = j
				break
			} else {
				last = nums[j+1]
			}
		}
	}
	return result
}
