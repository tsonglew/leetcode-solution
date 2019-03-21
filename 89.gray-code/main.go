package main

import "math"

func main() {

}

func grayCode(n int) []int {
	numCnt := int(math.Pow(2.0, float64(n)))
	result := make([]int, numCnt)
	for i := 0; i < numCnt; i++ {
		result[i] = binaryToGray(i)
	}
	return result
}

func binaryToGray(num int) int {
	return num ^ (num >> 1)
}
