package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {}

func reverse(s []string) []string {
	reversed := s
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - i - 1
		reversed[i], reversed[opp] = s[opp], s[i]
	}
	return reversed
}

func makeNZeros(n int) []string {
	zeros := make([]string, 0)
	for i := 0; i < n; i++ {
		zeros = append(zeros, "0")
	}
	return zeros
}

func addBinary(a string, b string) string {
	positions := make([]string, 0)
	aLen := len(a)
	bLen := len(b)
	length := int(math.Max(float64(aLen), float64(bLen)))
	anums := strings.Split(a, "")
	bnums := strings.Split(b, "")

	aReversed := reverse(anums)
	bReversed := reverse(bnums)

	aCompleted := append(aReversed, makeNZeros(length-aLen)...)
	bCompleted := append(bReversed, makeNZeros(length-bLen)...)

	f := 0

	for i := range aCompleted {
		aint, _ := strconv.Atoi(aCompleted[i])
		bint, _ := strconv.Atoi(bCompleted[i])
		sum := aint + bint + f
		if sum == 0 {
			positions = append(positions, "0")
			f = 0
		} else if sum == 1 {
			positions = append(positions, "1")
			f = 0
		} else if sum == 2 {
			positions = append(positions, "0")
			f = 1
		} else if sum == 3 {
			positions = append(positions, "1")
			f = 1
		}
	}

	if f == 1 {
		positions = append(positions, "1")
	}

	return strings.Join(reverse(positions), "")
}
