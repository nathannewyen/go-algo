package matrix

// TransposeMatrix swaps rows and columns of the matrix
func TransposeMatrix(grid [][]int) [][]int {
	rowCount := len(grid)
	columnCount := len(grid[0])

	// Create new matrix with swapped dimensions
	transposedGrid := make([][]int, columnCount)
	for columnIndex := range transposedGrid {
		transposedGrid[columnIndex] = make([]int, rowCount)
	}

	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
			transposedGrid[columnIndex][rowIndex] = grid[rowIndex][columnIndex]
		}
	}
	return transposedGrid
}
