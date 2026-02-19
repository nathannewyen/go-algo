package search

import "golang.org/x/exp/constraints"

// BinarySearch finds the index of a target value in a sorted slice.
// Returns the index and true if found, -1 and false otherwise.
// The input slice must be sorted in ascending order for correct results.
func BinarySearch[T constraints.Ordered](sortedItems []T, targetValue T) (int, bool) {
	lowIndex := 0
	highIndex := len(sortedItems) - 1

	for lowIndex <= highIndex {
		// Calculate midpoint without integer overflow
		middleIndex := lowIndex + (highIndex-lowIndex)/2

		if sortedItems[middleIndex] == targetValue {
			return middleIndex, true
		}

		if sortedItems[middleIndex] < targetValue {
			lowIndex = middleIndex + 1
		} else {
			highIndex = middleIndex - 1
		}
	}

	return -1, false
}

// BinarySearchFirstOccurrence finds the index of the first occurrence
// of the target value in a sorted slice. Useful when duplicates exist.
// Returns -1 and false if the value is not found.
func BinarySearchFirstOccurrence[T constraints.Ordered](sortedItems []T, targetValue T) (int, bool) {
	lowIndex := 0
	highIndex := len(sortedItems) - 1
	resultIndex := -1

	for lowIndex <= highIndex {
		middleIndex := lowIndex + (highIndex-lowIndex)/2

		if sortedItems[middleIndex] == targetValue {
			// Found a match, but keep searching left for an earlier occurrence
			resultIndex = middleIndex
			highIndex = middleIndex - 1
		} else if sortedItems[middleIndex] < targetValue {
			lowIndex = middleIndex + 1
		} else {
			highIndex = middleIndex - 1
		}
	}

	if resultIndex == -1 {
		return -1, false
	}
	return resultIndex, true
}

// TwoSum finds two indices in a sorted slice whose values add up to the target.
// Uses the two-pointer technique for O(n) time complexity.
// Returns the indices and true if found, (-1, -1) and false otherwise.
func TwoSum[T constraints.Integer](sortedItems []T, targetSum T) (int, int, bool) {
	leftPointer := 0
	rightPointer := len(sortedItems) - 1

	for leftPointer < rightPointer {
		currentSum := sortedItems[leftPointer] + sortedItems[rightPointer]

		if currentSum == targetSum {
			return leftPointer, rightPointer, true
		}

		// If the sum is too small, move the left pointer right to increase it
		if currentSum < targetSum {
			leftPointer++
		} else {
			// If the sum is too large, move the right pointer left to decrease it
			rightPointer--
		}
	}

	return -1, -1, false
}
