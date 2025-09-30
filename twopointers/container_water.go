package twopointers

// MaxWaterContainer finds two lines that form container holding most water
func MaxWaterContainer(heights []int) int {
	leftPointer := 0
	rightPointer := len(heights) - 1
	maximumArea := 0

	for leftPointer < rightPointer {
		// Calculate area between the two pointers
		containerWidth := rightPointer - leftPointer
		containerHeight := min(heights[leftPointer], heights[rightPointer])
		currentArea := containerWidth * containerHeight

		if currentArea > maximumArea {
			maximumArea = currentArea
		}

		// Move the pointer with the shorter height inward
		if heights[leftPointer] < heights[rightPointer] {
			leftPointer++
		} else {
			rightPointer--
		}
	}
	return maximumArea
}

// min returns the smaller of two integers
func min(firstValue int, secondValue int) int {
	if firstValue < secondValue {
		return firstValue
	}
	return secondValue
}
