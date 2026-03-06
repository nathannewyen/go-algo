package matrix

// SearchSortedMatrix searches for target in matrix where rows and columns are sorted
// Uses staircase search: O(m + n) time complexity
func SearchSortedMatrix(grid [][]int, targetValue int) bool {
	if len(grid) == 0 {
		return false
	}

	rowCount := len(grid)
	columnCount := len(grid[0])
	// Start from top-right corner
	currentRow := 0
	currentColumn := columnCount - 1

	for currentRow < rowCount && currentColumn >= 0 {
		currentElement := grid[currentRow][currentColumn]
		if currentElement == targetValue {
			return true
		} else if currentElement > targetValue {
			// Target must be in a column to the left
			currentColumn--
		} else {
			// Target must be in a row below
			currentRow++
		}
	}
	return false
}
