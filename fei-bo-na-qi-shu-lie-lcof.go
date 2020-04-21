func fib(n int) int {
    return Fib(n, 0, 1)
}

func Fib(n, a, b int) int {
    if n == 0 {
        return a
    }
    return Fib(n-1, b%1000000007, (a+b)%1000000007)
}
