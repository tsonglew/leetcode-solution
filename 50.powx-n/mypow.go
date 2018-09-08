package mypow

func myPow(x float64, n int) float64 {
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n%2 == 0 {
		p := myPow(x, n/2)
		return p * p
	} else {
		p := myPow(x, (n-1)/2)
		return x * p * p
	}
}
