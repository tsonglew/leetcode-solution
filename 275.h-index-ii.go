func hIndex(citations []int) int {
	mx := 0
	for i := len(citations)-1; i >=0; i-- {
		if citations[i] > len(citations)-i-1 {
			mx = len(citations)-i
		}
	}
	return mx
}
