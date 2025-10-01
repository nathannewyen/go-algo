package backtracking

// GenerateSubsets returns all possible subsets (power set) of input numbers
func GenerateSubsets(numbers []int) [][]int {
	allSubsets := [][]int{}
	currentSubset := []int{}
	buildSubsets(numbers, 0, currentSubset, &allSubsets)
	return allSubsets
}

// buildSubsets recursively includes or excludes each element
func buildSubsets(numbers []int, decisionIndex int, currentSubset []int, allSubsets *[][]int) {
	// Add current subset state as a valid subset
	subsetCopy := make([]int, len(currentSubset))
	copy(subsetCopy, currentSubset)
	*allSubsets = append(*allSubsets, subsetCopy)

	for elementIndex := decisionIndex; elementIndex < len(numbers); elementIndex++ {
		// Include this element in the subset
		currentSubset = append(currentSubset, numbers[elementIndex])
		buildSubsets(numbers, elementIndex+1, currentSubset, allSubsets)
		// Exclude this element (backtrack)
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
}
