func rangeBitwiseAnd(m int, n int) int {
    var i int
    for i = 0; m != n; i++ {
        m, n = m>>1, n>>1
    }
    return n << uint(i)
}