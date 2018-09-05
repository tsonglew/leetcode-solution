class Solution {
    public void solveSudoku(char[][] board) {
        	solveSubSudo(board);
	}

	public boolean solveSubSudo(char[][] board) {
		for (int i = 0; i < 9; ++i) {
			for (int j = 0; j < 9; ++j) {
				if (board[i][j] == '.') {
					for (char c = '1'; c <= '9'; ++c) {
						if (validSudoku(board, i, j, c)) {
                            board[i][j] = c;
							boolean f = solveSubSudo(board);
                            if (!f) {
                                board[i][j] = '.';
                            } else if (f) {
                                return true;
                            }
						}
					}
					return false;
				}
			}
		}
        return true;
	}

	public boolean validSudoku(char[][] board, int i, int j, char c) {
		int rowPre = 3 * (i / 3);
		int colPre = 3 * (j / 3);
		for (int m = 0; m < 9; ++m) {
			if (m!=j && board[i][m]==c) {
				return false;
			}
			if (m!=i && board[m][j]==c) {
				return false;
			}
			int rowIdx = rowPre + m / 3;
			int colIdx = colPre + m % 3;
			if (rowIdx!=i && colIdx!=j && board[rowIdx][colIdx]==c) {
				return false;
			}
		}
        return true;
    }
}