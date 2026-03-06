package matrix

// RotateClockwise90 rotates an NxN matrix 90 degrees clockwise in-place
func RotateClockwise90(grid [][]int) {
	matrixSize := len(grid)

	// Step 1: Transpose the matrix (swap rows and columns)
	for rowIndex := 0; rowIndex < matrixSize; rowIndex++ {
		for columnIndex := rowIndex + 1; columnIndex < matrixSize; columnIndex++ {
			grid[rowIndex][columnIndex], grid[columnIndex][rowIndex] = grid[columnIndex][rowIndex], grid[rowIndex][columnIndex]
		}
	}

	// Step 2: Reverse each row
	for rowIndex := 0; rowIndex < matrixSize; rowIndex++ {
		leftPointer := 0
		rightPointer := matrixSize - 1
		for leftPointer < rightPointer {
			grid[rowIndex][leftPointer], grid[rowIndex][rightPointer] = grid[rowIndex][rightPointer], grid[rowIndex][leftPointer]
			leftPointer++
			rightPointer--
		}
	}
}

// RotateCounterClockwise90 rotates an NxN matrix 90 degrees counter-clockwise
func RotateCounterClockwise90(grid [][]int) {
	matrixSize := len(grid)

	// Step 1: Transpose the matrix
	for rowIndex := 0; rowIndex < matrixSize; rowIndex++ {
		for columnIndex := rowIndex + 1; columnIndex < matrixSize; columnIndex++ {
			grid[rowIndex][columnIndex], grid[columnIndex][rowIndex] = grid[columnIndex][rowIndex], grid[rowIndex][columnIndex]
		}
	}

	// Step 2: Reverse each column
	for columnIndex := 0; columnIndex < matrixSize; columnIndex++ {
		topPointer := 0
		bottomPointer := matrixSize - 1
		for topPointer < bottomPointer {
			grid[topPointer][columnIndex], grid[bottomPointer][columnIndex] = grid[bottomPointer][columnIndex], grid[topPointer][columnIndex]
			topPointer++
			bottomPointer--
		}
	}
}
