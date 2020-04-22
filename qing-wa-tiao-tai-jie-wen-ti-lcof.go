func numWays(n int) int {
    a, b := 1, 1
    for ;n > 0; n-- {
        a, b = b, (b+a)%1000000007
    }
    return a
}
