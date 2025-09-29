package sorting

// MergeSort divides array in half, sorts each half, then merges
// Time: O(n log n), Space: O(n)
func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	middleIndex := len(numbers) / 2
	leftHalf := MergeSort(numbers[:middleIndex])
	rightHalf := MergeSort(numbers[middleIndex:])
	return mergeSortedHalves(leftHalf, rightHalf)
}

// mergeSortedHalves combines two sorted slices into one sorted slice
func mergeSortedHalves(leftHalf []int, rightHalf []int) []int {
	mergedResult := make([]int, 0, len(leftHalf)+len(rightHalf))
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(leftHalf) && rightIndex < len(rightHalf) {
		if leftHalf[leftIndex] <= rightHalf[rightIndex] {
			mergedResult = append(mergedResult, leftHalf[leftIndex])
			leftIndex++
		} else {
			mergedResult = append(mergedResult, rightHalf[rightIndex])
			rightIndex++
		}
	}
	// Append remaining elements from either half
	mergedResult = append(mergedResult, leftHalf[leftIndex:]...)
	mergedResult = append(mergedResult, rightHalf[rightIndex:]...)
	return mergedResult
}
