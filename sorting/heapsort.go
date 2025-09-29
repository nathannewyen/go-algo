package sorting

// HeapSort sorts by building a max heap then extracting elements
// Time: O(n log n), Space: O(1)
func HeapSort(numbers []int) []int {
	totalElements := len(numbers)

	// Build max heap from unordered array
	for parentIndex := totalElements/2 - 1; parentIndex >= 0; parentIndex-- {
		heapifySubtree(numbers, totalElements, parentIndex)
	}

	// Extract max element and place at end, then re-heapify
	for lastIndex := totalElements - 1; lastIndex > 0; lastIndex-- {
		numbers[0], numbers[lastIndex] = numbers[lastIndex], numbers[0]
		heapifySubtree(numbers, lastIndex, 0)
	}
	return numbers
}

// heapifySubtree ensures max-heap property for subtree rooted at rootIndex
func heapifySubtree(numbers []int, heapSize int, rootIndex int) {
	largestIndex := rootIndex
	leftChildIndex := 2*rootIndex + 1
	rightChildIndex := 2*rootIndex + 2

	if leftChildIndex < heapSize && numbers[leftChildIndex] > numbers[largestIndex] {
		largestIndex = leftChildIndex
	}
	if rightChildIndex < heapSize && numbers[rightChildIndex] > numbers[largestIndex] {
		largestIndex = rightChildIndex
	}

	// If largest is not root, swap and continue heapifying
	if largestIndex != rootIndex {
		numbers[rootIndex], numbers[largestIndex] = numbers[largestIndex], numbers[rootIndex]
		heapifySubtree(numbers, heapSize, largestIndex)
	}
}
