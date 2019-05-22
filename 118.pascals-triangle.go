func generate(numRows int) [][]int {
	result := [][]int{}
	if numRows == 0 {
		return result
	}
	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}
	result = append(result, []int{1, 1})
	if numRows == 2 {
		return result
	}
	for rows := numRows - 2; rows > 0; rows-- {
		newRow := []int{1}
		lastRow := result[len(result)-1]
		for i := 1; i < len(lastRow); i++ {
			newRow = append(newRow, lastRow[i-1]+lastRow[i])
		}
		newRow = append(newRow, 1)
		result = append(result, newRow)
	}
	return result
}