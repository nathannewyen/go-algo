package monotonic

// SlidingWindowMaximum finds the maximum in each sliding window of given size
func SlidingWindowMaximum(numbers []int, windowSize int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	maximums := []int{}
	// Deque stores indices of potentially maximum elements in decreasing order
	candidateIndices := []int{}

	for currentIndex := 0; currentIndex < len(numbers); currentIndex++ {
		// Remove indices that have fallen outside the current window
		for len(candidateIndices) > 0 && candidateIndices[0] <= currentIndex-windowSize {
			candidateIndices = candidateIndices[1:]
		}
		// Remove indices whose values are smaller than the current element
		for len(candidateIndices) > 0 && numbers[candidateIndices[len(candidateIndices)-1]] < numbers[currentIndex] {
			candidateIndices = candidateIndices[:len(candidateIndices)-1]
		}
		candidateIndices = append(candidateIndices, currentIndex)
		// Once we have processed at least windowSize elements, record the maximum
		if currentIndex >= windowSize-1 {
			maximums = append(maximums, numbers[candidateIndices[0]])
		}
	}
	return maximums
}
