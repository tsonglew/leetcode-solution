package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{[]byte{'1', '0', '1', '1', '0', '1', '1'}}))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	mxcnt := 0
	for i := range grid {
		for j := range grid[i] {
			if visited[i][j] {
				continue
			}
			if grid[i][j] == '1' {
				mxcnt++
				spread(i, j, grid, visited)
			}
		}
	}
	return mxcnt
}

func spread(i, j int, grid [][]byte, visited [][]bool) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || visited[i][j] || grid[i][j] == '0' {
		return
	}
	if grid[i][j] == '1' {
		visited[i][j] = true
	}
	spread(i-1, j, grid, visited)
	spread(i, j-1, grid, visited)
	spread(i+1, j, grid, visited)
	spread(i, j+1, grid, visited)
}
