package recursion

// WordSearch checks if a word exists in the grid by traversing adjacent cells
func WordSearch(grid [][]byte, targetWord string) bool {
	rowCount := len(grid)
	if rowCount == 0 {
		return false
	}
	columnCount := len(grid[0])
	// Try starting the search from every cell in the grid
	for row := 0; row < rowCount; row++ {
		for col := 0; col < columnCount; col++ {
			if searchFromCell(grid, targetWord, row, col, 0, rowCount, columnCount) {
				return true
			}
		}
	}
	return false
}

// searchFromCell performs DFS from a cell to match remaining characters
func searchFromCell(grid [][]byte, targetWord string, row int, col int, charIndex int, rowCount int, columnCount int) bool {
	if charIndex == len(targetWord) {
		return true
	}
	if row < 0 || row >= rowCount || col < 0 || col >= columnCount {
		return false
	}
	if grid[row][col] != targetWord[charIndex] {
		return false
	}
	// Mark cell as visited by temporarily changing its value
	originalChar := grid[row][col]
	grid[row][col] = '#'
	// Explore all four adjacent directions
	directionOffsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, offset := range directionOffsets {
		nextRow := row + offset[0]
		nextCol := col + offset[1]
		if searchFromCell(grid, targetWord, nextRow, nextCol, charIndex+1, rowCount, columnCount) {
			grid[row][col] = originalChar
			return true
		}
	}
	grid[row][col] = originalChar
	return false
}
