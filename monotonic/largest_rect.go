package monotonic

// LargestRectangleInHistogram finds the area of largest rectangle in a histogram
func LargestRectangleInHistogram(barHeights []int) int {
	barCount := len(barHeights)
	// Stack stores indices of bars in increasing height order
	heightStack := []int{}
	maxRectangleArea := 0

	for currentBar := 0; currentBar <= barCount; currentBar++ {
		currentHeight := 0
		if currentBar < barCount {
			currentHeight = barHeights[currentBar]
		}
		// Pop bars taller than current and calculate their max rectangle
		for len(heightStack) > 0 && barHeights[heightStack[len(heightStack)-1]] > currentHeight {
			poppedBarIndex := heightStack[len(heightStack)-1]
			heightStack = heightStack[:len(heightStack)-1]
			poppedHeight := barHeights[poppedBarIndex]
			// Width extends from just after the new stack top to current position
			rectangleWidth := currentBar
			if len(heightStack) > 0 {
				rectangleWidth = currentBar - heightStack[len(heightStack)-1] - 1
			}
			rectangleArea := poppedHeight * rectangleWidth
			if rectangleArea > maxRectangleArea {
				maxRectangleArea = rectangleArea
			}
		}
		heightStack = append(heightStack, currentBar)
	}
	return maxRectangleArea
}
