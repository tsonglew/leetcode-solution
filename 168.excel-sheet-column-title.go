func convertToTitle(n int) string {
    m := map[int]string{}
    for i := 0; i < 26; i++ {
        m[i] = string(i+65)
    }
    n -= 1
    result := ""
    for n >= 0 {
        result = m[n%26] + result
        n = n / 26 - 1
    }
    return result
}