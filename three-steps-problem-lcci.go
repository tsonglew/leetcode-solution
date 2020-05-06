func waysToStep(n int) int {
    b, c, d := 1, 2, 4
    switch n {
        case 0: return 1
        case 1: return 1
        case 2: return 2
        case 3: return 4
    }
    for i := 3; i < n; i++ {
        b, c, d = c, d, (b+c+d)%1000000007
    }
    return d
}
