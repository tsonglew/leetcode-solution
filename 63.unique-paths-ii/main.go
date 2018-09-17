package main

func main() {
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	grid[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && obstacleGrid[i-1][j] == 0 {
				grid[i][j] += grid[i-1][j]
			}
			if j > 0 && obstacleGrid[i][j-1] == 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[m-1][n-1]
}
