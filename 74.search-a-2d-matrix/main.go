package main

func main() {

}

func getCor(position, rowLen, colLen int) (int, int) {
	col := position % colLen
	row := position / colLen
	return col, row
}

func searchMatrix(matrix [][]int, target int) bool {
	rowLen := len(matrix)
	if rowLen == 0 {
		return false
	}
	colLen := len(matrix[0])
	if colLen == 0 {
		return false
	}
	low, high := 0, rowLen*colLen-1
	for low <= high {
		mid := (low + high) / 2
		col, row := getCor(mid, rowLen, colLen)
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			low = mid + 1
		} else if matrix[row][col] > target {
			high = mid - 1
		}
	}
	return false
}
