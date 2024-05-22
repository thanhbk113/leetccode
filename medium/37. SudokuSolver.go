package main

func solveSudoku(board [][]byte) {
	checkRow, checkCol, checkSquare := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				checkRow[r][board[r][c]-'0'-1] = true
				checkCol[c][board[r][c]-'0'-1] = true
				checkSquare[getSquareIndex(r, c)][board[r][c]-'0'-1] = true
			}
		}
	}

	backtrackSudoku(board, 0, 0, checkCol, checkRow, checkSquare)
}

func backtrackSudoku(board [][]byte, row, col int, checkCol, checkRow, checkSquare [9][9]bool) [][]byte {
	if col == 9 {
		row++
		col = 0
	}

	if row == 9 {
		return board
	}

	if board[row][col] != '.' {
		return backtrackSudoku(board, row, col+1, checkCol, checkRow, checkSquare)
	}

	if board[row][col] == '.' {

		for i := 1; i <= 9; i++ {
			if !checkCol[col][i-1] && !checkRow[row][i-1] && !checkSquare[getSquareIndex(row, col)][i-1] {
				board[row][col] = byte(i) + '0'
				checkCol[col][board[row][col]-'0'-1] = true
				checkRow[row][board[row][col]-'0'-1] = true
				checkSquare[getSquareIndex(row, col)][board[row][col]-'0'-1] = true
				if v := backtrackSudoku(board, row, col+1, checkCol, checkRow, checkSquare); v != nil {
					return v
				}
				checkCol[col][board[row][col]-'0'-1] = false
				checkRow[row][board[row][col]-'0'-1] = false
				checkSquare[getSquareIndex(row, col)][board[row][col]-'0'-1] = false
				board[row][col] = '.'
			}
		}
	}

	return nil
}

func getSquareIndex(r, c int) int {
	return r/3*3 + c/3
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
}
