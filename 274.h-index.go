func hIndex(citations []int) int {
	mx := 0
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i := range citations {
		if citations[i] >= i+1 {
			mx = i+1
		}
	}
	return mx
}
