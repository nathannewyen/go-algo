package sorting

// SelectionSort finds minimum element and places it at the beginning
// Time: O(n^2), Space: O(1)
func SelectionSort(numbers []int) []int {
	totalElements := len(numbers)
	for currentPosition := 0; currentPosition < totalElements-1; currentPosition++ {
		minimumIndex := currentPosition
		// Find the index of the smallest remaining element
		for searchIndex := currentPosition + 1; searchIndex < totalElements; searchIndex++ {
			if numbers[searchIndex] < numbers[minimumIndex] {
				minimumIndex = searchIndex
			}
		}
		// Swap minimum element into current position
		numbers[currentPosition], numbers[minimumIndex] = numbers[minimumIndex], numbers[currentPosition]
	}
	return numbers
}
