func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    return searchTarget(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, target)
}

func searchTarget(matrix [][]int, rowL, rowH, colL, colH, target int) bool {
    if rowL > rowH || colL > colH {
        return false
    }
    midRow := int((rowL + rowH) / 2)
    midCol := int((colL + colH) / 2)
    mid := matrix[midRow][midCol]
    if mid == target {
        return true
    } else if mid > target {
        return searchTarget(matrix, rowL, midRow-1, colL, colH, target) ||
        searchTarget(matrix, rowL, rowH, colL, midCol-1, target)
    } else {
        return searchTarget(matrix, midRow+1, rowH, colL, colH, target) ||
        searchTarget(matrix, rowL, midRow, midCol+1, colH, target)
    }
}