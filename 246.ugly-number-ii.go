func nthUglyNumber(n int) int {
    uglys := make([]int, n)
    uglys[0] = 1
    c2, c3, c5 := 0, 0, 0
    for i := 1; i < n; i++ {
        uglys[i] = min(uglys[c2]*2, uglys[c3]*3, uglys[c5]*5)
        if uglys[i] == uglys[c2]*2 {
            c2++
        }
        if uglys[i] == uglys[c3]*3 {
            c3++
        }
        if uglys[i] == uglys[c5]*5 {
            c5++
        }
    }
    return uglys[n-1]
}

func min(nums ...int) int {
    mn := nums[0]
    for _, n := range nums {
        if n < mn {
            mn = n
        }
    }
    return mn
}
