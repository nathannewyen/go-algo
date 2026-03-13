package hashing

// IsValidSudoku checks if a 9x9 sudoku board is valid
func IsValidSudoku(board [][]byte) bool {
	// Track seen digits for each row, column, and 3x3 box
	rowSeen := [9]map[byte]bool{}
	colSeen := [9]map[byte]bool{}
	boxSeen := [9]map[byte]bool{}
	for i := 0; i < 9; i++ {
		rowSeen[i] = map[byte]bool{}
		colSeen[i] = map[byte]bool{}
		boxSeen[i] = map[byte]bool{}
	}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cellValue := board[row][col]
			if cellValue == '.' {
				continue
			}
			boxIndex := (row/3)*3 + col/3
			// Check if digit already exists in its row, column, or box
			if rowSeen[row][cellValue] {
				return false
			}
			if colSeen[col][cellValue] {
				return false
			}
			if boxSeen[boxIndex][cellValue] {
				return false
			}
			rowSeen[row][cellValue] = true
			colSeen[col][cellValue] = true
			boxSeen[boxIndex][cellValue] = true
		}
	}
	return true
}
