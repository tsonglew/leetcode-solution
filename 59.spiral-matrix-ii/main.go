package main

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	result := make([][]int, n)
	flag := RIGHT
	for i := range result {
		result[i] = make([]int, n)
	}
	i, j, cnt := 0, 0, 1

	for {
		if i >= n || i < 0 || j >= n || j < 0 || result[i][j] != 0 {
			break
		}
		result[i][j] = cnt
		cnt++
		if flag == RIGHT {
			if j == n-1 || result[i][j+1] != 0 {
				flag = DOWN
				i++
			} else {
				j++
			}
		} else if flag == DOWN {
			if i == n-1 || result[i+1][j] != 0 {
				flag = LEFT
				j--
			} else {
				i++
			}
		} else if flag == LEFT {
			if j == 0 || result[i][j-1] != 0 {
				flag = UP
				i--
			} else {
				j--
			}
		} else if flag == UP {
			if i == 0 || result[i-1][j] != 0 {
				flag = RIGHT
				j++
			} else {
				i--
			}
		}

	}
	return result
}
