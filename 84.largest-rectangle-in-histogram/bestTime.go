package main

func largestRectangleAreaTime(heights []int) int {
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
