package hashing

// ContainsDuplicate checks if any value appears at least twice in the array
func ContainsDuplicate(numbers []int) bool {
	seenNumbers := map[int]bool{}
	for _, currentNumber := range numbers {
		if seenNumbers[currentNumber] {
			return true
		}
		seenNumbers[currentNumber] = true
	}
	return false
}

// ContainsNearbyDuplicate checks if there are two equal elements within k distance
func ContainsNearbyDuplicate(numbers []int, maxDistance int) bool {
	// Sliding window set of size maxDistance
	recentValues := map[int]bool{}
	for currentIndex, currentNumber := range numbers {
		if recentValues[currentNumber] {
			return true
		}
		recentValues[currentNumber] = true
		// Shrink window when it exceeds maxDistance
		if currentIndex >= maxDistance {
			evictedValue := numbers[currentIndex-maxDistance]
			delete(recentValues, evictedValue)
		}
	}
	return false
}
