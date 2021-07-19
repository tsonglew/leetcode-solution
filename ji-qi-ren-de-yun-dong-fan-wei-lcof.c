int movingCount(int m, int n, int k){
    int** board = malloc(sizeof(int*) * m);
    for (int i=0; i<m;++i)
        board[i] = malloc(sizeof(int) * n);
    board[0][0] = 1;
    int res = 1;
    for (int i=0; i<m; ++i) {
        for (int j=0; j<n; ++j) {
            if (i==0&&j==0)
                continue;
            if ((i>0&&board[i-1][j] == 1||j>0&&board[i][j-1]==1)&&(i/10+i%10+j/10+j%10<=k)) {
                board[i][j] = 1;
                res++;
            } else {
                board[i][j] = 0;
            }
        }
    }
    return res;
}
