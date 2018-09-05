package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	resultInts := make([]int, len(num1)+len(num2))
	resultStrs := make([]string, len(num1)+len(num2))
	trailZeroFlag := true
	for j := len(num2) - 1; j >= 0; j-- {
		for i := len(num1) - 1; i >= 0; i-- {
			mul := (int(num1[i]) - 48) * (int(num2[j]) - 48)
			sum := resultInts[i+j+1] + mul
			resultInts[i+j+1] = sum % 10
			resultInts[i+j] += sum / 10
		}
	}
	for i := range resultInts {
		if trailZeroFlag && resultInts[i] == 0 {
			continue
		}
		trailZeroFlag = false
		resultStrs = append(resultStrs, strconv.Itoa(resultInts[i]))
	}
	if trailZeroFlag {
		resultStrs = append(resultStrs, "0")
	}
	return strings.Join(resultStrs, "")
}

func main() {
	fmt.Println(multiply("123", "456"))
}
