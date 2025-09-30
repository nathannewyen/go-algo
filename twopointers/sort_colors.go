package twopointers

// SortColors sorts array of 0s, 1s, and 2s in-place (Dutch National Flag)
func SortColors(colorArray []int) {
	// Three pointers: low for 0s boundary, high for 2s boundary
	lowBoundary := 0
	currentIndex := 0
	highBoundary := len(colorArray) - 1

	for currentIndex <= highBoundary {
		switch colorArray[currentIndex] {
		case 0:
			// Swap current element to low section
			colorArray[lowBoundary], colorArray[currentIndex] = colorArray[currentIndex], colorArray[lowBoundary]
			lowBoundary++
			currentIndex++
		case 1:
			// 1s stay in the middle
			currentIndex++
		case 2:
			// Swap current element to high section
			colorArray[currentIndex], colorArray[highBoundary] = colorArray[highBoundary], colorArray[currentIndex]
			highBoundary--
		}
	}
}
