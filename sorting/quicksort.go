package sorting

// QuickSort sorts using divide and conquer with a pivot element
// Time: O(n log n) average, O(n^2) worst, Space: O(log n)
func QuickSort(numbers []int) []int {
	quickSortRecursive(numbers, 0, len(numbers)-1)
	return numbers
}

// quickSortRecursive sorts the subarray between lowIndex and highIndex
func quickSortRecursive(numbers []int, lowIndex int, highIndex int) {
	if lowIndex < highIndex {
		pivotPosition := partitionArray(numbers, lowIndex, highIndex)
		quickSortRecursive(numbers, lowIndex, pivotPosition-1)
		quickSortRecursive(numbers, pivotPosition+1, highIndex)
	}
}

// partitionArray places pivot in correct position with smaller elements left
func partitionArray(numbers []int, lowIndex int, highIndex int) int {
	pivotValue := numbers[highIndex]
	// smallerElementBoundary tracks where smaller elements end
	smallerElementBoundary := lowIndex - 1

	for currentIndex := lowIndex; currentIndex < highIndex; currentIndex++ {
		if numbers[currentIndex] <= pivotValue {
			smallerElementBoundary++
			numbers[smallerElementBoundary], numbers[currentIndex] = numbers[currentIndex], numbers[smallerElementBoundary]
		}
	}
	numbers[smallerElementBoundary+1], numbers[highIndex] = numbers[highIndex], numbers[smallerElementBoundary+1]
	return smallerElementBoundary + 1
}
