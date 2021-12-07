func colorBorder(grid [][]int, row int, col int, color int) [][]int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return grid
    }
    visited := make([][]int, len(grid))
    for i := range visited {
        visited[i] = make([]int, len(grid[0]))
        for j := range visited[i] {
            visited[i][j] = 0
        }
    }
    search(grid, visited, row, col, color, grid[row][col])
    return grid
}

func search(grid, visited [][]int, row, col int, color, origin int) {
    if (visited[row][col] == 1) {
        return
    }
    if grid[row][col] != origin {
        return
    }
    visited[row][col] = 1

    if row == 0 || row == len(grid)-1 || col == 0 || col == len(grid[0])-1 {
        grid[row][col] = color
    }
    if valid(row-1, col, len(grid)-1, len(grid[0])-1, visited) && 
        grid[row-1][col] != grid[row][col] ||
        valid(row+1, col, len(grid)-1, len(grid[0])-1, visited) && 
        grid[row+1][col] != grid[row][col] ||
        valid(row, col-1, len(grid)-1, len(grid[0])-1, visited) && 
        grid[row][col-1] != grid[row][col] ||
        valid(row, col+1, len(grid)-1, len(grid[0])-1, visited) && 
        grid[row][col+1] != grid[row][col] {
            grid[row][col] = color
        }
    if valid(row-1, col, len(grid)-1, len(grid[0])-1, visited) {
        search(grid, visited, row-1, col, color, origin)
    }
    if valid(row+1, col, len(grid)-1, len(grid[0])-1, visited) {
        search(grid, visited, row+1, col, color, origin)
    }
    if valid(row, col-1, len(grid)-1, len(grid[0])-1, visited) {
        search(grid, visited, row, col-1, color, origin)
    }
    if valid(row, col+1, len(grid)-1, len(grid[0])-1, visited) {
        search(grid, visited, row, col+1, color, origin)
    }
}

func valid(row, col, rowMax, colMax int, visited [][]int) bool {
    if row < 0 || row > rowMax || col < 0 || col > colMax || visited[row][col] == 1 {
        return false
    }
    return true
}
