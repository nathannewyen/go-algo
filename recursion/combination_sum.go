package recursion

// CombinationSum finds all unique combinations that sum to the target
func CombinationSum(candidates []int, targetSum int) [][]int {
	validCombinations := [][]int{}
	currentCombination := []int{}
	findCombinations(candidates, targetSum, 0, currentCombination, &validCombinations)
	return validCombinations
}

// findCombinations recursively builds combinations that sum to remaining target
func findCombinations(candidates []int, remainingTarget int, startIndex int, currentCombination []int, validCombinations *[][]int) {
	if remainingTarget == 0 {
		combinationCopy := make([]int, len(currentCombination))
		copy(combinationCopy, currentCombination)
		*validCombinations = append(*validCombinations, combinationCopy)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		candidateValue := candidates[i]
		if candidateValue > remainingTarget {
			continue
		}
		currentCombination = append(currentCombination, candidateValue)
		// Allow reusing the same candidate by passing i instead of i+1
		findCombinations(candidates, remainingTarget-candidateValue, i, currentCombination, validCombinations)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}
