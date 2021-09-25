class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) == 0:
            return 0
        min_price = prices[0]
        profit = 0
        for i in prices[1:]:
            profit = max(profit, i-min_price)
            min_price = min(min_price, i)
        return profit
