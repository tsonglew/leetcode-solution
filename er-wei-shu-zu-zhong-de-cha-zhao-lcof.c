bool binSearchRow(int** matrix, int row, int colLow, int colHigh, int target) {
    int colMid;
    while (colLow <= colHigh) {
        colMid = (colLow+colHigh) / 2;
        if (matrix[row][colMid] == target) {
            return true;
        } else if (matrix[row][colMid] > target) {
            colHigh = colMid - 1;
        } else {
            colLow = colMid + 1;
        }
    }
    return false;
}

bool binSearchCol(int** matrix, int col, int rowLow, int rowHigh, int target) {
    int rowMid;
    while (rowLow <= rowHigh) {
        rowMid = (rowLow+rowHigh) /2;
        if (matrix[rowMid][col] == target) {
            return true;
        } else if (matrix[rowMid][col] > target) {
            rowHigh = rowMid - 1;
        } else {
            rowLow = rowMid + 1;
        }
    }
    return false;
}

bool find(int** matrix, int rowLow, int rowHigh, int colLow, int colHigh, int target) {
    if (rowLow <= rowHigh && colLow <= colHigh) {
        int rowMid = (rowLow + rowHigh) / 2;
        int colMid = (colLow + colHigh) / 2;
        if (binSearchRow(matrix, rowMid, colLow, colHigh, target))
            return true;
        if (binSearchCol(matrix, colMid, rowLow, rowHigh, target))
            return true;
        if (rowMid-1 >= rowLow && colMid-1>=colLow && matrix[rowLow][colLow] <= target &&  matrix[rowMid-1][colMid-1] >= target) {
            if (find(matrix, rowLow, rowMid-1, colLow, colMid-1, target)) {
                return true;
            }
        } 
        if (rowMid+1 <= rowHigh && colMid+1<=colHigh && matrix[rowHigh][colHigh] >= target &&  matrix[rowMid+1][colMid+1] <= target) {
            if (find(matrix, rowMid+1, rowHigh, colMid+1, colHigh, target)) {
                return true;
            }
        }
        if (colMid+1<=colHigh && rowMid-1>=rowLow && matrix[rowLow][colMid+1]<=target && matrix[rowMid-1][colHigh] >= target) {
            if (find(matrix, rowLow, rowMid-1, colMid+1, colHigh, target)) {
                return true;
            }
        }
        if (colMid-1>=colLow && rowMid+1<=rowHigh && matrix[rowHigh][colMid-1]>=target && matrix[rowMid+1][colLow] <= target) {
            if (find(matrix, rowMid + 1, rowHigh, colLow, colMid - 1, target)) {
                return true;
            }
        }
    }
    return false;

}

bool findNumberIn2DArray(int** matrix, int matrixSize, int* matrixColSize, int target){
    if (matrixSize == 0 || matrixColSize[0] == 0)
        return false;
    return find(matrix, 0, matrixSize-1, 0, matrixColSize[0]-1, target);
}
