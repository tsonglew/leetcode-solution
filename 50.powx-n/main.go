package main

import (
	"fmt"
)

func main(){
	fmt.Println(myPow(2.0000, 10))
}
func myPow(x float64, n int) float64 {
	m := map[int]float64{}
	return pow(x, n, m)
}

func pow(x float64, n int, m map[int]float64) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if value, ok := m[n]; ok {
		return value
	}
	if n < 0 {
		n = -n
		x = 1/x
	}
	if n % 2 == 0 {
		p := myPow(x, n/2)
		m[n/2] = p
		return p*p
	} else {
		p := myPow(x, (n-1)/2)
		m[(n-1)/2] = p
		return x * p * p
	}
}