package stringalgo

// ZAlgorithmSearch finds all occurrences of pattern in text using Z-array
func ZAlgorithmSearch(textContent string, searchPattern string) []int {
	// Concatenate pattern + separator + text
	concatenated := searchPattern + "$" + textContent
	zArray := buildZArray(concatenated)
	patternLength := len(searchPattern)

	matchPositions := []int{}
	for arrayIndex := patternLength + 1; arrayIndex < len(zArray); arrayIndex++ {
		if zArray[arrayIndex] == patternLength {
			// Match found, convert back to position in original text
			matchPositions = append(matchPositions, arrayIndex-patternLength-1)
		}
	}
	return matchPositions
}

// buildZArray computes Z-array where Z[i] = length of longest substring
// starting at position i that matches a prefix of the string
func buildZArray(inputString string) []int {
	stringLength := len(inputString)
	zValues := make([]int, stringLength)
	// [windowLeft, windowRight] tracks the rightmost Z-box
	windowLeft := 0
	windowRight := 0

	for currentIndex := 1; currentIndex < stringLength; currentIndex++ {
		if currentIndex < windowRight {
			// Use previously computed value as starting point
			correspondingIndex := currentIndex - windowLeft
			zValues[currentIndex] = zValues[correspondingIndex]
			if zValues[currentIndex] > windowRight-currentIndex {
				zValues[currentIndex] = windowRight - currentIndex
			}
		}
		// Extend Z-value by character comparison
		for currentIndex+zValues[currentIndex] < stringLength &&
			inputString[zValues[currentIndex]] == inputString[currentIndex+zValues[currentIndex]] {
			zValues[currentIndex]++
		}
		// Update Z-box if we extended past windowRight
		if currentIndex+zValues[currentIndex] > windowRight {
			windowLeft = currentIndex
			windowRight = currentIndex + zValues[currentIndex]
		}
	}
	return zValues
}
