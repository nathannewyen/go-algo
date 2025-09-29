package sorting

// CountingSort sorts integers by counting occurrences of each value
// Time: O(n + k) where k is the range of input, Space: O(k)
func CountingSort(numbers []int) []int {
	if len(numbers) == 0 {
		return numbers
	}

	// Find the maximum value to determine count array size
	maximumValue := numbers[0]
	for _, currentNumber := range numbers {
		if currentNumber > maximumValue {
			maximumValue = currentNumber
		}
	}

	// Count occurrences of each value
	occurrenceCount := make([]int, maximumValue+1)
	for _, currentNumber := range numbers {
		occurrenceCount[currentNumber]++
	}

	// Rebuild the sorted array from counts
	sortedIndex := 0
	for value, count := range occurrenceCount {
		for repeatIndex := 0; repeatIndex < count; repeatIndex++ {
			numbers[sortedIndex] = value
			sortedIndex++
		}
	}
	return numbers
}
