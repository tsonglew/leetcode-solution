func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrime[i] {
			for j := i; i*j < n; j++ {
				isPrime[i*j] = false
			}
		}
	}
	cnt := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
		}
	}	
	return cnt
}