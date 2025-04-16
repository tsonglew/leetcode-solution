class Solution {

    private boolean[][] visited;
    private char[][] board;
    private int m;
    private int n;
    private int[][] direction;
    private boolean[] canReachBoarder;

    public void solve(char[][] board) {
        this.board = board;
        this.m = board.length;
        this.n = board[0].length;
        this.visited = new boolean[this.m][this.n];
        this.canReachBoarder = new boolean[this.m*this.n];
        this.direction = new int[][]{{0,-1}, {0,1}, {1,0}, {-1,0}};

        for (int i = 0; i < this.m; i++) {
            this.dfs(i, 0);
            this.dfs(i, this.n-1);
        }
        for (int j = 0; j < this.n; j++) {
            this.dfs(0, j);
            this.dfs(this.m-1, j);
        }

        for (int i = 0; i < this.m; i++) {
            for (int j = 0; j < this.n; j++) {
                if (this.board[i][j] == 'X') {
                    continue;
                }
                if (!this.canReachBoarder[this.hash(i, j)]) {
                    this.board[i][j] = 'X';
                }
            }
        }
    }

    private void dfs(int i, int j) {
        if (!this.valid(i, j)) {
            return;
        }
        if (this.visited[i][j]) {
            return;
        }
        this.visited[i][j] = true;
        if (this.board[i][j] == 'X') {
            return;
        }
        this.canReachBoarder[this.hash(i,j)] = true;
        for (int[] d : this.direction) {
            this.dfs(i+d[0], j+d[1]);
        }
    }

    private boolean valid(int i, int j) {
        return i >= 0 && j >= 0 && i < this.m && j < this.n;
    }

    private int hash(int i, int j) {
        return i * this.n + j;
    }

}
