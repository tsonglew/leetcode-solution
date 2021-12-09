func validTicTacToe(board []string) bool {
    xCnt, oCnt := pieceCnt(board)
    if oCnt > xCnt || xCnt > oCnt + 1 {
        return false
    }
    xWin := pieceWin(board, 'X')
    oWin := pieceWin(board, 'O')
    if xWin && oWin {
        return false
    }
    if xWin && xCnt <= oCnt {
        return false
    }
    if oWin && xCnt > oCnt {
        return false
    }
    return true
}

func pieceCnt(board []string) (int, int) {
    xCnt, oCnt := 0, 0
    for i := range board {
        for j := 0; j < len(board); j++ {
            if board[i][j] == 'X' {
                xCnt ++
            } else if board[i][j] == 'O' {
                oCnt ++
            }
        }
    }
    return xCnt, oCnt
}

func pieceWin(board []string, piece byte) bool {
    if board[0][0] == piece &&  board[1][1] == piece &&  board[2][2] == piece || board[0][2] == piece &&  board[1][1] == piece &&  board[2][0] == piece {
        return true
    }
    for i := 0; i < len(board); i++ {
        if board[i] == string([]byte{piece, piece, piece}) {
            return true
        }
        if board[0][i] == piece && board[1][i] == piece && board[2][i] == piece {
            return true
        }
    }
    return false
}
