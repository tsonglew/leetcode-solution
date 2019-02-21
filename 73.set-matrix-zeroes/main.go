package main

func main() {

}

func setZeroes(matrix [][]int) {
	if len(matrix) <= 1 || len(matrix[0]) <= 1 {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[i][j] == 0 {
					for m := 0; m < len(matrix); m++ {
						matrix[m][j] = 0
					}
					for n := 0; n < len(matrix[0]); n++ {
						matrix[i][n] = 0
					}
					return
				}
			}
		}
		return
	}
	col, row := false, false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i != 0 && j != 0 {
					matrix[i][0], matrix[0][j] = 0, 0
				}
				if i == 0 {
					row = true
				}
				if j == 0 {
					col = true
				}
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for n := 1; n < len(matrix[0]); n++ {
				matrix[i][n] = 0
			}
		}
	}

	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for n := 1; n < len(matrix); n++ {
				matrix[n][i] = 0
			}
		}
	}

	if col {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if row {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}
