package main

func isValidSudoku(board [][]byte) bool {
	rowsMap := [9][9]bool{}
	colsMap := [9][9]bool{}
	squaresMap := [9][9]bool{}

	n := 9

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == '.' {
				continue
			}
			val := board[r][c] - '0' - 1
			squareIndex := r/3*3 + c/3
			if rowsMap[r][val] || colsMap[c][val] || squaresMap[squareIndex][val] {
				return false
			}
			rowsMap[r][val] = true
			colsMap[c][val] = true
			squaresMap[squareIndex][val] = true

		}

	}
	return true
}

func main() {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	isValidSudoku(board)
}
