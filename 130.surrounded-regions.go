package main

import "fmt"

type Point struct {
	i, j int
}

func solve(board [][]byte) {
	var rows, cols int
	if rows = len(board); rows == 0 {
		return
	}
	if cols = len(board[0]); cols == 0 {
		return
	}
	searched := map[Point]bool{}
	marked := map[Point]bool{}
	for i, j := 0, 0; j < cols; j++ {
		if board[i][j] == 'O' {
			deepSearch(board, i, j, rows, cols, searched, marked)
		}
	}
	for i, j := rows-1, 0; j < cols; j++ {
		if board[i][j] == 'O' {
			deepSearch(board, i, j, rows, cols, searched, marked)
		}
	}
	for i, j := 1, 0; i < rows-1; i++ {
		if board[i][j] == 'O' {
			deepSearch(board, i, j, rows, cols, searched, marked)
		}
	}
	for i, j := 1, cols-1; i < rows-1; i++ {
		if board[i][j] == 'O' {
			deepSearch(board, i, j, rows, cols, searched, marked)
		}
	}
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if board[i][j] == 'O' && !searched[Point{i, j}] {
				board[i][j] = 'X'
			}
		}
	}
}

func deepSearch(board [][]byte, i, j, rows, cols int, searched, marked map[Point]bool) {
	if b, ok := searched[Point{i, j}]; ok && b == true {
		return
	}
	searched[Point{i, j}] = true
	marked[Point{i, j}] = true
	if i > 0 && board[i-1][j] == 'O' {
		deepSearch(board, i-1, j, rows, cols, searched, marked)
	}
	if i < rows-1 && board[i+1][j] == 'O' {
		deepSearch(board, i+1, j, rows, cols, searched, marked)
	}
	if j > 0 && board[i][j-1] == 'O' {
		deepSearch(board, i, j-1, rows, cols, searched, marked)
	}
	if j < cols-1 && board[i][j+1] == 'O' {
		deepSearch(board, i, j+1, rows, cols, searched, marked)
	}
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board)
}
