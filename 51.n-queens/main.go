package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	r := [][]string{}
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

func subSolveNQueens(current [][]int, depth, n int, r *[][]string) {
	if depth >= n {
		result := []string{}
		for i := range current {
			tmp := []string{}
			for j := range current[i] {
				if current[i][j] == 0 {
					tmp = append(tmp, ".")
				} else if current[i][j] == 1 {
					tmp = append(tmp, "Q")
				}
			}
			result = append(result, strings.Join(tmp, ""))
		}
		*r = append(*r, result)
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

func main() {
	fmt.Println(solveNQueens(5))
}
