import "math"

func maximalSquare(matrix [][]byte) int {
	mx := 0
	areas := make([][]int, len(matrix))
	for i := range areas {
		areas[i] = make([]int, len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == '1' {
				areas[i][j] = 1
				if i > 0 && j > 0 {
					if int(areas[i-1][j-1]) > 0 {
						l := int(math.Sqrt(float64(areas[i-1][j-1])))
						width := getWidth(matrix, i, j, l)
						areas[i][j] = width * width
					}
				}
				if areas[i][j] > mx {
					mx = areas[i][j]
				}
			} else {
				areas[i][j] = 0
			}
		}
	}
	return mx
}

func getWidth(matrix [][]byte, i, j, max int) int {
	for l := 0; l < max; l++ {
		if matrix[i][j-l-1] == '0' || matrix[i-l-1][j] == '0' {
			return l + 1
		}
	}
	return max + 1
}