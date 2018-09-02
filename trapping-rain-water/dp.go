package main

import (
	"fmt"
	"math"
	"time"
)

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(height[i])))
		ri := len(height) - i - 1
		rightMax[ri] = int(math.Max(float64(rightMax[ri+1]), float64(height[ri])))
	}
	total := 0
	for i := range leftMax {
		toAdd := int(math.Min(float64(leftMax[i]), float64(rightMax[i]))) - height[i]
		total += toAdd
	}
	return total
}

func main() {
	start := time.Now()
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("costed: %v", time.Since(start))
}
