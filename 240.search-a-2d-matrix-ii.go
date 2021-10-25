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

func searchMatrixMethod2(matrix [][]int, target int) bool {
    if len(matrix) > 0 {
        return f(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, target)
    }
    return false
}

func f(matrix [][]int, i1, i2, j1, j2, target int) bool {
    if i1 > i2 || j1 > j2 || i1 < 0 || i2 >= len(matrix) || j1 < 0 || j2 > len(matrix[0]) {
        return false
    }
    mid_i, mid_j := (i1+i2)/2, (j1+j2)/2
    mid := matrix[mid_i][mid_j]
    if mid == target {
        return true
    } else if mid < target {
        return f(matrix, mid_i+1, i2, j1, j2, target) || f(matrix, i1, mid_i, mid_j+1, j2, target)
    }
    return f(matrix, i1, mid_i-1, j1, j2, target) || f(matrix, mid_i, i2, j1, mid_j-1, target)
}
