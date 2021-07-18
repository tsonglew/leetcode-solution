bool search(char** board, int boardRowSize, int boardColSize, char* word, int i, int j) {
    if (word[0] == '\0')
        return true;
    if (i < 0 || i >= boardRowSize || j < 0 || j >= boardColSize) 
        return false;
    if (board[i][j] == word[0]) {
        char origin = board[i][j];
        board[i][j] = ' ';
        bool res = search(board, boardRowSize, boardColSize, word+1, i-1, j) || search(board, boardRowSize, boardColSize, word+1, i+1, j) || search(board, boardRowSize, boardColSize, word+1, i, j-1) || search(board, boardRowSize, boardColSize, word+1, i, j+1);
        board[i][j] = origin;
        return res;
    }
        
    return false;
}


bool exist(char** board, int boardSize, int* boardColSize, char* word){
    for (int i=0; i < boardSize; ++i) {
        for (int j=0; j < boardColSize[0]; ++j) {
            if (search(board, boardSize, boardColSize[0], word, i, j))
                return true;
        }
    }
    return false;
}

