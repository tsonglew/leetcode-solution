class Solution {

    int m;
    int n;
    int[][] obstacleGrid;
    int[][] routes;

    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        this.obstacleGrid = obstacleGrid;
        this.m = obstacleGrid.length;
        this.n = obstacleGrid[0].length;
        this.routes = new int[this.m][this.n];
        
        for (int i = 0; i < this.m; i++) {
            for (int j = 0; j < this.n; j++) {
                this.routes[i][j] = -1;
            }
        }
        return this.calc(0 ,0);

    }

    public int calc(int i, int j) {

        if (!this.isPositionValid(i, j)) {
            return 0;
        }

        if (this.routes[i][j] == -1) {

            if (this.hasObstacle(i, j)) {
                this.routes[i][j] = 0;
            } else {
                int val = (i == this.m - 1 && j == this.n - 1) ? 1 : 0;
                val += this.calc(i+1, j);
                val += this.calc(i, j+1);
                this.routes[i][j] = val;
            }
        }

        return this.routes[i][j];
    }

    public boolean isPositionValid(int i, int j) {
        return i >= 0 && i < this.m && j >= 0 && j < this.n;
    }

    public boolean hasObstacle(int i, int j) {
        return this.isPositionValid(i, j) && this.obstacleGrid[i][j] == 1;
    }
}
