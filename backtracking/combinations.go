package backtracking

// GenerateCombinations returns all combinations of k elements from 1..n
func GenerateCombinations(totalRange int, combinationSize int) [][]int {
	allCombinations := [][]int{}
	currentCombination := []int{}
	buildCombinations(1, totalRange, combinationSize, currentCombination, &allCombinations)
	return allCombinations
}

// buildCombinations recursively builds combinations starting from startValue
func buildCombinations(startValue int, totalRange int, remainingSize int, currentCombination []int, allCombinations *[][]int) {
	if remainingSize == 0 {
		completedCombination := make([]int, len(currentCombination))
		copy(completedCombination, currentCombination)
		*allCombinations = append(*allCombinations, completedCombination)
		return
	}

	for candidateValue := startValue; candidateValue <= totalRange; candidateValue++ {
		currentCombination = append(currentCombination, candidateValue)
		buildCombinations(candidateValue+1, totalRange, remainingSize-1, currentCombination, allCombinations)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}
