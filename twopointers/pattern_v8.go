package twopointers

// SquareSortedArray8 computes squares of sorted array maintaining order
func SquareSortedArray8(sortedNumbers []int) []int {
	resultLength := len(sortedNumbers)
	squaredResult := make([]int, resultLength)
	leftPointer := 0
	rightPointer := resultLength - 1
	insertPosition := resultLength - 1

	for leftPointer <= rightPointer {
		leftSquare := sortedNumbers[leftPointer] * sortedNumbers[leftPointer]
		rightSquare := sortedNumbers[rightPointer] * sortedNumbers[rightPointer]
		if leftSquare > rightSquare {
			squaredResult[insertPosition] = leftSquare
			leftPointer++
		} else {
			squaredResult[insertPosition] = rightSquare
			rightPointer--
		}
		insertPosition--
	}
	return squaredResult
}
