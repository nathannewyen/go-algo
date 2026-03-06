package matrix

// MaxRectangleInHistogram finds largest rectangle area in histogram bars
func MaxRectangleInHistogram(barHeights []int) int {
	// Stack stores indices of bars in increasing height order
	heightStack := []int{}
	maximumArea := 0

	for currentIndex := 0; currentIndex <= len(barHeights); currentIndex++ {
		currentHeight := 0
		if currentIndex < len(barHeights) {
			currentHeight = barHeights[currentIndex]
		}

		for len(heightStack) > 0 && currentHeight < barHeights[heightStack[len(heightStack)-1]] {
			// Pop the top bar and calculate its rectangle area
			topBarIndex := heightStack[len(heightStack)-1]
			heightStack = heightStack[:len(heightStack)-1]

			rectangleHeight := barHeights[topBarIndex]
			rectangleWidth := currentIndex
			if len(heightStack) > 0 {
				rectangleWidth = currentIndex - heightStack[len(heightStack)-1] - 1
			}

			rectangleArea := rectangleHeight * rectangleWidth
			if rectangleArea > maximumArea {
				maximumArea = rectangleArea
			}
		}
		heightStack = append(heightStack, currentIndex)
	}
	return maximumArea
}

// MaxRectangleInMatrix finds largest rectangle containing only 1s in binary matrix
func MaxRectangleInMatrix(binaryGrid [][]int) int {
	if len(binaryGrid) == 0 {
		return 0
	}

	columnCount := len(binaryGrid[0])
	cumulativeHeights := make([]int, columnCount)
	maximumRectangle := 0

	for rowIndex := 0; rowIndex < len(binaryGrid); rowIndex++ {
		// Build histogram heights row by row
		for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
			if binaryGrid[rowIndex][columnIndex] == 1 {
				cumulativeHeights[columnIndex]++
			} else {
				cumulativeHeights[columnIndex] = 0
			}
		}
		// Find max rectangle in current histogram
		currentMaxArea := MaxRectangleInHistogram(cumulativeHeights)
		if currentMaxArea > maximumRectangle {
			maximumRectangle = currentMaxArea
		}
	}
	return maximumRectangle
}
