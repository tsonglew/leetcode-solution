func maxAreaOfIsland(grid [][]int) int {
    ans := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                area := 1
                dfs(grid, i, j, &area)
                if area > ans { ans = area }
            }
        }
    }
    return ans
}

func dfs(grid [][]int, i, j int, area *int) {
    grid[i][j] = 0
    if i + 1 < len(grid) && grid[i+1][j] == 1 {
        *area ++
        dfs(grid, i+1, j, area)
    }
    if i > 0 && grid[i-1][j] == 1 {
        *area ++
        dfs(grid, i-1, j, area)
    }
    if j + 1 < len(grid[0]) && grid[i][j+1] == 1 {
        *area ++
        dfs(grid, i, j+1, area)
    }
    if j > 0 && grid[i][j-1] == 1 {
        *area ++
        dfs(grid, i, j-1, area)
    }
}
