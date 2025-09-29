package sorting

// BubbleSort sorts a slice by repeatedly swapping adjacent elements
// Time: O(n^2), Space: O(1)
func BubbleSort(numbers []int) []int {
	totalElements := len(numbers)
	for outerIndex := 0; outerIndex < totalElements-1; outerIndex++ {
		swapOccurred := false
		for innerIndex := 0; innerIndex < totalElements-outerIndex-1; innerIndex++ {
			if numbers[innerIndex] > numbers[innerIndex+1] {
				numbers[innerIndex], numbers[innerIndex+1] = numbers[innerIndex+1], numbers[innerIndex]
				swapOccurred = true
			}
		}
		// Early termination if no swaps occurred
		if !swapOccurred {
			break
		}
	}
	return numbers
}
