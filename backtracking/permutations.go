package backtracking

// GeneratePermutations returns all possible orderings of the input numbers
func GeneratePermutations(numbers []int) [][]int {
	allPermutations := [][]int{}
	currentPermutation := []int{}
	usedIndices := make([]bool, len(numbers))
	buildPermutations(numbers, currentPermutation, usedIndices, &allPermutations)
	return allPermutations
}

// buildPermutations recursively constructs permutations by choosing unused elements
func buildPermutations(numbers []int, currentPermutation []int, usedIndices []bool, allPermutations *[][]int) {
	// Base case: permutation is complete
	if len(currentPermutation) == len(numbers) {
		completedPermutation := make([]int, len(currentPermutation))
		copy(completedPermutation, currentPermutation)
		*allPermutations = append(*allPermutations, completedPermutation)
		return
	}

	for elementIndex := 0; elementIndex < len(numbers); elementIndex++ {
		if usedIndices[elementIndex] {
			continue
		}
		// Choose the element
		usedIndices[elementIndex] = true
		currentPermutation = append(currentPermutation, numbers[elementIndex])

		buildPermutations(numbers, currentPermutation, usedIndices, allPermutations)

		// Backtrack: unchoose the element
		currentPermutation = currentPermutation[:len(currentPermutation)-1]
		usedIndices[elementIndex] = false
	}
}
