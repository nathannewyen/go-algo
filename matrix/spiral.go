package matrix

// SpiralOrder returns elements of matrix in spiral order
func SpiralOrder(grid [][]int) []int {
	if len(grid) == 0 {
		return []int{}
	}

	spiralResult := []int{}
	topBoundary := 0
	bottomBoundary := len(grid) - 1
	leftBoundary := 0
	rightBoundary := len(grid[0]) - 1

	for topBoundary <= bottomBoundary && leftBoundary <= rightBoundary {
		// Traverse right along top boundary
		for columnIndex := leftBoundary; columnIndex <= rightBoundary; columnIndex++ {
			spiralResult = append(spiralResult, grid[topBoundary][columnIndex])
		}
		topBoundary++

		// Traverse down along right boundary
		for rowIndex := topBoundary; rowIndex <= bottomBoundary; rowIndex++ {
			spiralResult = append(spiralResult, grid[rowIndex][rightBoundary])
		}
		rightBoundary--

		// Traverse left along bottom boundary
		if topBoundary <= bottomBoundary {
			for columnIndex := rightBoundary; columnIndex >= leftBoundary; columnIndex-- {
				spiralResult = append(spiralResult, grid[bottomBoundary][columnIndex])
			}
			bottomBoundary--
		}

		// Traverse up along left boundary
		if leftBoundary <= rightBoundary {
			for rowIndex := bottomBoundary; rowIndex >= topBoundary; rowIndex-- {
				spiralResult = append(spiralResult, grid[rowIndex][leftBoundary])
			}
			leftBoundary++
		}
	}
	return spiralResult
}
