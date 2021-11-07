func maxCount(m int, n int, ops [][]int) int {
    maxM, maxN := m, n
    for _, op := range ops {
        if op[0] * op[1] == 0 {
            continue
        }
        maxM = min(maxM, op[0])
        maxN = min(maxN, op[1])
    }
    return maxM * maxN
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
