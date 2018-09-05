import java.util.HashSet;

class Solution {
    public boolean isValidSudoku(char[][] board) {
        for (int i = 0; i < board.length; ++i) {
            HashSet<Character> rows = new HashSet<>();
            HashSet<Character> cols = new HashSet<>();
            HashSet<Character> cubs = new HashSet<>();
            for (int j = 0; j < board[i].length; ++j) {
                if (board[i][j] != '.' && !rows.add(board[i][j])) {
                    return false;
                }
                if (board[j][i] != '.' && !cols.add(board[j][i])) {
                    return false;
                }
                int rowIdx = 3 * (i / 3) + j/3 ;
                int colIdx = 3 * (i % 3) + j%3;
                if (board[rowIdx][colIdx] != '.' && !cubs.add(board[rowIdx][colIdx])) {
                    return false;
                }
            }
        }
        return true;
    }
}
