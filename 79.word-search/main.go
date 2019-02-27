package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	if len(board) < 1 || len(board[0]) < 1 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				used := make(map[point]bool)
				used[point{i, j}] = true
				if search(i, j, 0, board, word, used) {
					return true
				}
			}
		}
	}
	return false
}

func search(m, n, p int, board [][]byte, word string, used map[point]bool) bool {
	if p == len(word)-1 {
		return true
	}
	for _, pt := range []point{
		{m - 1, n},
		{m + 1, n},
		{m, n - 1},
		{m, n + 1},
	} {
		if ud, ok := used[pt]; pt.x >= 0 && pt.x < len(board) &&
			pt.y >= 0 && pt.y < len(board[0]) &&
			(!ok || ud == false) &&
			board[pt.x][pt.y] == word[p+1] {
			newUsed := make(map[point]bool)
			copyMap(newUsed, used)
			newUsed[pt] = true
			if search(pt.x, pt.y, p+1, board, word, newUsed) {
				return true
			}
		}
	}
	return false
}

func copyMap(m1, m2 map[point]bool) {
	for key, value := range m2 {
		m1[key] = value
	}
}
