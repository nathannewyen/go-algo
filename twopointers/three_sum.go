package twopointers

import "sort"

// FindThreeNumbersWithTargetSum finds all unique triplets that sum to target
func FindThreeNumbersWithTargetSum(numbers []int, targetSum int) [][]int {
	sort.Ints(numbers)
	foundTriplets := [][]int{}

	for firstIndex := 0; firstIndex < len(numbers)-2; firstIndex++ {
		// Skip duplicate values for the first element
		if firstIndex > 0 && numbers[firstIndex] == numbers[firstIndex-1] {
			continue
		}

		leftPointer := firstIndex + 1
		rightPointer := len(numbers) - 1
		remainingTarget := targetSum - numbers[firstIndex]

		for leftPointer < rightPointer {
			twoElementSum := numbers[leftPointer] + numbers[rightPointer]
			if twoElementSum == remainingTarget {
				foundTriplets = append(foundTriplets, []int{numbers[firstIndex], numbers[leftPointer], numbers[rightPointer]})
				leftPointer++
				rightPointer--
				// Skip duplicates for second element
				for leftPointer < rightPointer && numbers[leftPointer] == numbers[leftPointer-1] {
					leftPointer++
				}
			} else if twoElementSum < remainingTarget {
				leftPointer++
			} else {
				rightPointer--
			}
		}
	}
	return foundTriplets
}
