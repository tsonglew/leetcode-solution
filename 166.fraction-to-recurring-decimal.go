package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(7, 6))
	fmt.Println(fractionToDecimal(1, 6))
	fmt.Println(fractionToDecimal(1, 13))
	fmt.Println(fractionToDecimal(-50, 8))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func fractionToDecimal(numerator int, denominator int) string {
	calculated := make(map[int]int)
	n1, n2 := abs(numerator), abs(denominator)
	result := strconv.Itoa(n1 / n2)
	if numerator*denominator < 0 {
		result = "-" + result
	}
	result2 := ""
	n1 = (n1 % n2) * 10
	for idx := 0; n1 != 0; idx++ {
		if v, ok := calculated[n1]; ok {
			result2 = result2[0:v] + "(" + result2[v:] + ")"
			break
		}
		calculated[n1] = idx
		if n1 < n2 {
			n1 *= 10
			result2 += "0"
			continue
		}
		value := strconv.Itoa(n1 / n2)
		result2 += value
		n1 = (n1 % n2) * 10
	}
	if len(result2) > 0 {
		result += "." + result2
	}
	return result
}
