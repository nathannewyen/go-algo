package matrix

// SetMatrixZeroes sets entire row and column to zero if element is zero
// Uses first row and column as markers for O(1) extra space
func SetMatrixZeroes(grid [][]int) {
	rowCount := len(grid)
	columnCount := len(grid[0])
	firstRowHasZero := false
	firstColumnHasZero := false

	// Check if first row contains zero
	for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
		if grid[0][columnIndex] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// Check if first column contains zero
	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		if grid[rowIndex][0] == 0 {
			firstColumnHasZero = true
			break
		}
	}

	// Use first row/column as markers for remaining cells
	for rowIndex := 1; rowIndex < rowCount; rowIndex++ {
		for columnIndex := 1; columnIndex < columnCount; columnIndex++ {
			if grid[rowIndex][columnIndex] == 0 {
				grid[rowIndex][0] = 0
				grid[0][columnIndex] = 0
			}
		}
	}

	// Zero out cells based on markers
	for rowIndex := 1; rowIndex < rowCount; rowIndex++ {
		for columnIndex := 1; columnIndex < columnCount; columnIndex++ {
			if grid[rowIndex][0] == 0 || grid[0][columnIndex] == 0 {
				grid[rowIndex][columnIndex] = 0
			}
		}
	}

	// Handle first row and column
	if firstRowHasZero {
		for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
			grid[0][columnIndex] = 0
		}
	}
	if firstColumnHasZero {
		for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
			grid[rowIndex][0] = 0
		}
	}
}
