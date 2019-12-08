func titleToNumber(s string) int {
    result := 0
    base := 1
    for i := range s {
        result += (int(s[len(s)-i-1])-64) * base
        base *= 26
    }
    return result
}