package totalnqueens

func totalNQueens(n int) int {
	r := 0
	for i := 0; i < n; i++ {
		board := make([][]int, n)
		for i := range board {
			board[i] = make([]int, n)
		}
		board[0][i] = 1
		subSolveNQueens(board, 1, n, &r)
	}
	return r
}

func subSolveNQueens(current [][]int, depth, n int, r *int) {
	if depth >= n {
		(*r)++
		return
	}
	for i := 0; i < n; i++ {
		if check(current, depth, i, n) {
			current[depth][i] = 1
			subSolveNQueens(current, depth+1, n, r)
			current[depth][i] = 0
		}
	}
}

func check(current [][]int, x, y, n int) bool {
	for i := 0; i < x; i++ {
		for j := 0; j < n; j++ {
			if float64(x-i)/float64(y-j) == 1 || float64(x-i)/float64(y-j) == -1 || j == y {
				if current[i][j] == 1 {
					return false
				}
			}
		}
	}
	return true
}
