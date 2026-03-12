package recursion

// SolveNQueens returns all valid N-Queens board configurations
func SolveNQueens(boardSize int) [][]string {
	allSolutions := [][]string{}
	// Track which columns and diagonals are under attack
	columnOccupied := make([]bool, boardSize)
	mainDiagonalOccupied := make([]bool, 2*boardSize)
	antiDiagonalOccupied := make([]bool, 2*boardSize)
	queenPositions := make([]int, boardSize)
	placeQueens(boardSize, 0, queenPositions, columnOccupied, mainDiagonalOccupied, antiDiagonalOccupied, &allSolutions)
	return allSolutions
}

// placeQueens tries placing a queen in each column of the current row
func placeQueens(boardSize int, currentRow int, queenPositions []int, columnOccupied []bool, mainDiagonalOccupied []bool, antiDiagonalOccupied []bool, allSolutions *[][]string) {
	if currentRow == boardSize {
		board := buildBoard(boardSize, queenPositions)
		*allSolutions = append(*allSolutions, board)
		return
	}
	for col := 0; col < boardSize; col++ {
		mainDiagIndex := currentRow - col + boardSize
		antiDiagIndex := currentRow + col
		if columnOccupied[col] || mainDiagonalOccupied[mainDiagIndex] || antiDiagonalOccupied[antiDiagIndex] {
			continue
		}
		queenPositions[currentRow] = col
		columnOccupied[col] = true
		mainDiagonalOccupied[mainDiagIndex] = true
		antiDiagonalOccupied[antiDiagIndex] = true
		placeQueens(boardSize, currentRow+1, queenPositions, columnOccupied, mainDiagonalOccupied, antiDiagonalOccupied, allSolutions)
		columnOccupied[col] = false
		mainDiagonalOccupied[mainDiagIndex] = false
		antiDiagonalOccupied[antiDiagIndex] = false
	}
}

// buildBoard creates a visual board representation from queen column positions
func buildBoard(boardSize int, queenPositions []int) []string {
	board := make([]string, boardSize)
	for row := 0; row < boardSize; row++ {
		rowChars := make([]byte, boardSize)
		for col := 0; col < boardSize; col++ {
			rowChars[col] = '.'
		}
		rowChars[queenPositions[row]] = 'Q'
		board[row] = string(rowChars)
	}
	return board
}
