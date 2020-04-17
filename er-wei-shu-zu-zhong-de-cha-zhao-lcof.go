func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    for i := 0; i < len(matrix); i++ {
        if matrix[i][0] > target {
            continue
        }
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == target {
                return true
            }
        }
    }
    return false
}
