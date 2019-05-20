package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{[]byte{'1', '0', '1', '1', '0', '1', '1'}}))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	mxcnt := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				mxcnt++
				spread(i, j, grid)
			}
		}
	}
	return mxcnt
}

func spread(i, j int, grid [][]byte) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	spread(i-1, j, grid)
	spread(i, j-1, grid)
	spread(i+1, j, grid)
	spread(i, j+1, grid)
}
