func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    buys := make([]int, len(prices))
    sell := make([]int, len(prices))
    cool := make([]int, len(prices))
    buys[0], sell[0], cool[0] = -prices[0], 0, 0
    last := len(prices)-1
    for i := 1; i <= last; i++ {
        buys[i] = max(buys[i-1], cool[i-1]-prices[i])
        sell[i] = max(buys[i-1]+prices[i], sell[i-1])
        cool[i] = max(cool[i-1], sell[i-1])
    }
    return max(buys[last], sell[last], cool[last])
}

func max(nums ...int) int {
    mx := nums[0]
    for i := range nums {
        if nums[i] > mx {
            mx = nums[i]
        }
    }
    return mx
}