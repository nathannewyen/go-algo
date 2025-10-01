package backtracking

// SearchWordInGrid checks if a word exists in the grid by moving adjacent cells
func SearchWordInGrid(characterGrid [][]byte, targetWord string) bool {
	rowCount := len(characterGrid)
	columnCount := len(characterGrid[0])

	for startRow := 0; startRow < rowCount; startRow++ {
		for startColumn := 0; startColumn < columnCount; startColumn++ {
			if exploreWordPath(characterGrid, targetWord, startRow, startColumn, 0) {
				return true
			}
		}
	}
	return false
}

// exploreWordPath recursively searches for remaining characters from current position
func exploreWordPath(characterGrid [][]byte, targetWord string, currentRow int, currentColumn int, characterIndex int) bool {
	if characterIndex == len(targetWord) {
		return true
	}
	// Boundary and character match checks
	if currentRow < 0 || currentRow >= len(characterGrid) || currentColumn < 0 || currentColumn >= len(characterGrid[0]) {
		return false
	}
	if characterGrid[currentRow][currentColumn] != targetWord[characterIndex] {
		return false
	}

	// Mark cell as visited by replacing character
	originalCharacter := characterGrid[currentRow][currentColumn]
	characterGrid[currentRow][currentColumn] = #

	// Explore all four directions
	directionOffsets := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, offset := range directionOffsets {
		if exploreWordPath(characterGrid, targetWord, currentRow+offset[0], currentColumn+offset[1], characterIndex+1) {
			characterGrid[currentRow][currentColumn] = originalCharacter
			return true
		}
	}

	// Backtrack: restore original character
	characterGrid[currentRow][currentColumn] = originalCharacter
	return false
}
