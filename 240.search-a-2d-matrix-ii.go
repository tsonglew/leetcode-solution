func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return false }
    for i := range matrix {
        if matrix[i][len(matrix[0])-1] < target {
            continue
        }
        if matrix[i][0] > target {
            return false
        }
        for j := range matrix[i] {
            if matrix[i][j] == target {
                return true
            } else if matrix[i][j] > target {
                break
            }
        }
    }
    return false
}
