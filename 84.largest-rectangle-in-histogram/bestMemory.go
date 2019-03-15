package main

func largestRectangleAreaMemory(heights []int) int {
	area := 0
	for i := range heights {
		l, h := i, i
		for l >= 0 && heights[l] >= heights[i] {
			l--
		}
		for h <= len(heights)-1 && heights[h] >= heights[i] {
			h++
		}
		current := (h - l - 1) * heights[i]
		if current > area {
			area = current
		}
	}
	return area
}
