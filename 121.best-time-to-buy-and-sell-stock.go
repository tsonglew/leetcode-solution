func maxProfit(prices []int) int {
	result := 0
	currentMinIndex := 0
	for i := 1; i < len(prices); i++ {
		n := prices[i] - prices[currentMinIndex]
		if n > result {
			result = n
		}
		if prices[currentMinIndex] > prices[i] {
			currentMinIndex = i
		}
	}
	return result
}