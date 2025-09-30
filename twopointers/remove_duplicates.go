package twopointers

// RemoveDuplicates removes duplicate values from sorted array in-place
// Returns the count of unique elements
func RemoveDuplicates(sortedNumbers []int) int {
	if len(sortedNumbers) == 0 {
		return 0
	}

	// uniqueInsertPosition tracks where the next unique element should go
	uniqueInsertPosition := 1

	for currentIndex := 1; currentIndex < len(sortedNumbers); currentIndex++ {
		// Only keep elements that differ from the previous unique element
		if sortedNumbers[currentIndex] != sortedNumbers[uniqueInsertPosition-1] {
			sortedNumbers[uniqueInsertPosition] = sortedNumbers[currentIndex]
			uniqueInsertPosition++
		}
	}
	return uniqueInsertPosition
}
