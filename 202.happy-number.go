func isHappy(n int) bool {
    n1, n2 := n, n
    for n1 != 0 && n2 != 0 {
        n1 = calcHappy(n1)
        n2 = calcHappy(calcHappy(n2))
        if n1 == 1 || n2 == 1 {
            return true
        }
        if n1 == n2 {
            return false
        }
    }
    return false
}

func calcHappy(n int) int {
    result := 0
    for n > 0 {
        b := n % 10
        result += b*b
        n /= 10
    }
    return result
}