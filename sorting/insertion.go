package sorting

// InsertionSort builds sorted array one element at a time
// Time: O(n^2), Space: O(1)
func InsertionSort(numbers []int) []int {
	for currentIndex := 1; currentIndex < len(numbers); currentIndex++ {
		currentValue := numbers[currentIndex]
		// Shift elements greater than currentValue to the right
		shiftIndex := currentIndex - 1
		for shiftIndex >= 0 && numbers[shiftIndex] > currentValue {
			numbers[shiftIndex+1] = numbers[shiftIndex]
			shiftIndex--
		}
		numbers[shiftIndex+1] = currentValue
	}
	return numbers
}
