package matrix

// DiagonalTraversal returns elements traversed diagonally
func DiagonalTraversal(grid [][]int) []int {
	if len(grid) == 0 {
		return []int{}
	}

	rowCount := len(grid)
	columnCount := len(grid[0])
	diagonalResult := []int{}

	// Total number of diagonals is rowCount + columnCount - 1
	totalDiagonals := rowCount + columnCount - 1

	for diagonalIndex := 0; diagonalIndex < totalDiagonals; diagonalIndex++ {
		if diagonalIndex%2 == 0 {
			// Traverse upward: start from bottom-left of diagonal
			startRow := diagonalIndex
			if startRow >= rowCount {
				startRow = rowCount - 1
			}
			startColumn := diagonalIndex - startRow

			for currentRow, currentColumn := startRow, startColumn; currentRow >= 0 && currentColumn < columnCount; {
				diagonalResult = append(diagonalResult, grid[currentRow][currentColumn])
				currentRow--
				currentColumn++
			}
		} else {
			// Traverse downward: start from top-right of diagonal
			startColumn := diagonalIndex
			if startColumn >= columnCount {
				startColumn = columnCount - 1
			}
			startRow := diagonalIndex - startColumn

			for currentRow, currentColumn := startRow, startColumn; currentRow < rowCount && currentColumn >= 0; {
				diagonalResult = append(diagonalResult, grid[currentRow][currentColumn])
				currentRow++
				currentColumn--
			}
		}
	}
	return diagonalResult
}
