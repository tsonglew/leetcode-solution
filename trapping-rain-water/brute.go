package main

import (
	"fmt"
	"math"
	"time"
)

func trap(height []int) int {
	total := 0
	if len(height) <= 2 {
		return total
	}
	for i := range height {
		maxLeft := height[i]
		maxRight := height[i]
		for j := i; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		for j := i; j < len(height); j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		border := math.Min(float64(maxLeft), float64(maxRight))
		total += (int(border) - height[i])
	}
	return total
}

func main() {
	start := time.Now()
	fmt.Println(trapBrute([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("costed: %v", time.Since(start))
}
