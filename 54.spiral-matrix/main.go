package main

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

func main(){}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) <=0 || len(matrix[0]) <= 0 {
		return []int{}
	}
	result := make([]int, len(matrix)*len(matrix[0]))
	flag := RIGHT
	m := make([][]int, len(matrix))
	for i := range m {
		m[i] = make([]int, len(matrix[0]))
	}
	i, j, cnt := 0, 0, 0

	for {
		if i>=len(matrix) || i<0 || j>=len(matrix[0]) || j<0|| m[i][j] == 1 {
			break
		}
		result[cnt] = matrix[i][j]
		cnt++
		m[i][j] = 1
		if flag == RIGHT {
			if j == len(matrix[0])-1 || m[i][j+1] == 1 {
				flag = DOWN
				i++
			} else {
				j++
			}
		} else if flag == DOWN {
			if i == len(matrix)-1 || m[i+1][j] == 1 {
				flag = LEFT
				j--
			} else {
				i++
			}
		} else if flag == LEFT {
			if j == 0 || m[i][j-1] == 1 {
				flag = UP
				i--
			} else {
				j--
			}
		} else if flag == UP {
			if i == 0 || m[i-1][j] == 1 {
				flag = RIGHT
				j++
			} else {
				i--
			}
		}

	}
	return result
}