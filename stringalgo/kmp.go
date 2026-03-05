package stringalgo

// KMPSearch finds all occurrences of pattern in text using Knuth-Morris-Pratt
func KMPSearch(textContent string, searchPattern string) []int {
	matchPositions := []int{}
	failureTable := buildKMPFailureTable(searchPattern)

	textIndex := 0
	patternIndex := 0

	for textIndex < len(textContent) {
		if textContent[textIndex] == searchPattern[patternIndex] {
			textIndex++
			patternIndex++
		}

		if patternIndex == len(searchPattern) {
			// Full match found at this position
			matchPositions = append(matchPositions, textIndex-patternIndex)
			patternIndex = failureTable[patternIndex-1]
		} else if textIndex < len(textContent) && textContent[textIndex] != searchPattern[patternIndex] {
			if patternIndex != 0 {
				// Use failure table to skip unnecessary comparisons
				patternIndex = failureTable[patternIndex-1]
			} else {
				textIndex++
			}
		}
	}
	return matchPositions
}

// buildKMPFailureTable computes longest proper prefix-suffix array
func buildKMPFailureTable(searchPattern string) []int {
	failureTable := make([]int, len(searchPattern))
	prefixLength := 0
	tableIndex := 1

	for tableIndex < len(searchPattern) {
		if searchPattern[tableIndex] == searchPattern[prefixLength] {
			prefixLength++
			failureTable[tableIndex] = prefixLength
			tableIndex++
		} else {
			if prefixLength != 0 {
				prefixLength = failureTable[prefixLength-1]
			} else {
				failureTable[tableIndex] = 0
				tableIndex++
			}
		}
	}
	return failureTable
}
