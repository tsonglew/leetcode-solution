func superPow(a int, b []int) int {
    mod := 1337 
    ans := quickPow(a, b[0],mod)
    for i := 1; i < len(b); i++ {
        ans = quickPow(ans,10,mod) * quickPow(a, b[i],mod) % mod
    }
    return ans
}

func quickPow(a, b, mod int) int {
    res := 1
    for b > 0 {
        if b & 1 == 1 {
            res = res * a % mod
        }    
        b >>= 1
        a = a * a % mod
    }
    return res
}
