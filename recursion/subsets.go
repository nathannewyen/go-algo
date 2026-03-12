package recursion

// GenerateSubsets returns all possible subsets of the input numbers
func GenerateSubsets(numbers []int) [][]int {
	allSubsets := [][]int{}
	currentSubset := []int{}
	buildSubsets(numbers, 0, currentSubset, &allSubsets)
	return allSubsets
}

// buildSubsets recursively includes or excludes each element
func buildSubsets(numbers []int, startIndex int, currentSubset []int, allSubsets *[][]int) {
	// Add copy of current subset to results at every recursion level
	subsetCopy := make([]int, len(currentSubset))
	copy(subsetCopy, currentSubset)
	*allSubsets = append(*allSubsets, subsetCopy)

	for i := startIndex; i < len(numbers); i++ {
		currentSubset = append(currentSubset, numbers[i])
		buildSubsets(numbers, i+1, currentSubset, allSubsets)
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
}
