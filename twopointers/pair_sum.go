package twopointers

// FindPairWithTargetSum finds two numbers in a sorted array that add up to target
// Returns indices of the pair, or [-1, -1] if not found
func FindPairWithTargetSum(sortedNumbers []int, targetSum int) [2]int {
	leftPointer := 0
	rightPointer := len(sortedNumbers) - 1

	for leftPointer < rightPointer {
		currentSum := sortedNumbers[leftPointer] + sortedNumbers[rightPointer]
		if currentSum == targetSum {
			return [2]int{leftPointer, rightPointer}
		} else if currentSum < targetSum {
			leftPointer++
		} else {
			rightPointer--
		}
	}
	return [2]int{-1, -1}
}
