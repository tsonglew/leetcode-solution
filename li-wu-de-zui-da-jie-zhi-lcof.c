int max(int i, int j) {
    if (i > j) {
        return i;
    }
    return j;
}

int maxValue(int** grid, int gridSize, int* gridColSize){
    for (int i = 0; i < gridSize; i++) {
        for (int j = 0; j < gridColSize[0]; j++) {
            if (i == 0 && j == 0) {
                continue;
            }
            if (i == 0) {
                grid[i][j] += grid[i][j-1];
                continue;
            }
            if (j == 0) {
                grid[i][j] += grid[i-1][j];
                continue;
            }
            grid[i][j] += max(grid[i][j-1], grid[i-1][j]);
        }
    }
    return grid[gridSize-1][gridColSize[0]-1];
}
