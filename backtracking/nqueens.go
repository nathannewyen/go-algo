package backtracking

// SolveNQueens finds all valid placements of n queens on an n x n board
func SolveNQueens(boardSize int) [][]int {
	allSolutions := [][]int{}
	// queenColumnPositions[row] stores the column where the queen is placed
	queenColumnPositions := make([]int, boardSize)
	placeQueens(boardSize, 0, queenColumnPositions, &allSolutions)
	return allSolutions
}

// placeQueens tries to place a queen in each column of the current row
func placeQueens(boardSize int, currentRow int, queenColumnPositions []int, allSolutions *[][]int) {
	if currentRow == boardSize {
		solution := make([]int, boardSize)
		copy(solution, queenColumnPositions)
		*allSolutions = append(*allSolutions, solution)
		return
	}

	for candidateColumn := 0; candidateColumn < boardSize; candidateColumn++ {
		if isQueenPlacementSafe(queenColumnPositions, currentRow, candidateColumn) {
			queenColumnPositions[currentRow] = candidateColumn
			placeQueens(boardSize, currentRow+1, queenColumnPositions, allSolutions)
		}
	}
}

// isQueenPlacementSafe checks if placing a queen at (row, col) conflicts with existing queens
func isQueenPlacementSafe(queenColumnPositions []int, targetRow int, targetColumn int) bool {
	for previousRow := 0; previousRow < targetRow; previousRow++ {
		previousColumn := queenColumnPositions[previousRow]
		// Check same column or diagonal conflicts
		if previousColumn == targetColumn {
			return false
		}
		rowDifference := targetRow - previousRow
		columnDifference := targetColumn - previousColumn
		if columnDifference < 0 {
			columnDifference = -columnDifference
		}
		if rowDifference == columnDifference {
			return false
		}
	}
	return true
}
