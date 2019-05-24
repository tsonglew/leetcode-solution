func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	lastSellPrice := prices[len(prices)-1]
	lastBuyPrice := lastSellPrice
	profit := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] >= lastBuyPrice {
			profit += (lastSellPrice - lastBuyPrice)
			lastSellPrice = prices[i]
			lastBuyPrice = lastSellPrice
		} else {
			lastBuyPrice = prices[i]
		}
	}
	if lastBuyPrice < lastSellPrice {
		profit += (lastSellPrice - lastBuyPrice)
	}
	return profit
}