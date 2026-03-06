package matrix

// ReshapeMatrix converts matrix to new dimensions if element count matches
func ReshapeMatrix(originalGrid [][]int, newRowCount int, newColumnCount int) [][]int {
	originalRowCount := len(originalGrid)
	originalColumnCount := len(originalGrid[0])
	totalElements := originalRowCount * originalColumnCount

	// Cannot reshape if total elements dont match
	if totalElements != newRowCount*newColumnCount {
		return originalGrid
	}

	reshapedGrid := make([][]int, newRowCount)
	for rowIndex := range reshapedGrid {
		reshapedGrid[rowIndex] = make([]int, newColumnCount)
	}

	// Map each element from old position to new position
	for elementIndex := 0; elementIndex < totalElements; elementIndex++ {
		oldRow := elementIndex / originalColumnCount
		oldColumn := elementIndex % originalColumnCount
		newRow := elementIndex / newColumnCount
		newColumn := elementIndex % newColumnCount
		reshapedGrid[newRow][newColumn] = originalGrid[oldRow][oldColumn]
	}
	return reshapedGrid
}
