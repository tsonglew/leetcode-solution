package main

import "fmt"

func main() {
	fmt.Println(maximalRectangle([][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '0', '1', '0'},
	}))
}

func maximalRectangle(matrix [][]byte) int {
	hs := make([][]int, len(matrix))
	la := 0
	for i := 0; i < len(matrix); i++ {
		hs[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 {
				hs[i][j] = int(matrix[i][j] - 48)
			} else {
				if matrix[i][j] == '0' {
					hs[i][j] = 0
				} else {
					hs[i][j] = hs[i-1][j] + 1
				}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		area := largestRectangleArea(hs[i])
		if area > la {
			la = area
		}
	}
	return la
}

// imported from 84
func largestRectangleArea(heights []int) int {
	area := 0
	if len(heights) == 0 {
		return area
	}
	le, re := make([]int, len(heights)), make([]int, len(heights))
	le[0], re[len(heights)-1] = -1, len(heights)
	for i := range heights {
		le[i], re[i] = i-1, i+1
	}
	for i := range heights {
		for l := i; l >= 0 && heights[l] >= heights[i]; l = le[l] {
			le[i] = le[l]
		}
	}
	for i := len(heights) - 1; i >= 0; i-- {
		for r := i; r < len(heights) && heights[r] >= heights[i]; r = re[r] {
			re[i] = re[r]
		}
	}

	for i := range heights {
		current := (re[i] - le[i] - 1) * heights[i]
		if current > area {
			area = current
		}
	}
	return area
}
