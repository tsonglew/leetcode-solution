package main

import "math"

func main() {
}

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			left, up := math.MaxInt64, math.MaxInt64
			if i > 0 {
				up = grid[i-1][j]
			}
			if j > 0 {
				left = grid[i][j-1]
			}
			grid[i][j] += mn(left, up)
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func mn(a, b int) int {
	if a < b {
		return a
	}
	return b
}
