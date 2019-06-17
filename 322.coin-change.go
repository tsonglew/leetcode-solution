func coinChange(coins []int, amount int) int {
    amounts := make([]int, amount+1)
    for i := 1; i < len(amounts); i++ {
        mn := -1
        for j := range coins {
            minused := i - coins[j]
            if minused == 0 {
                mn = 1
                break
            } else if minused > 0 && amounts[minused] > 0 && (mn < 0 || (mn > 0  && amounts[minused] + 1 < mn)) {
                    mn = amounts[minused] + 1
                } 
        }
        amounts[i] = mn
    }
    return amounts[amount]
}