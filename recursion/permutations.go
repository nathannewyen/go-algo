package recursion

// GeneratePermutations returns all possible orderings of the input numbers
func GeneratePermutations(numbers []int) [][]int {
	allPermutations := [][]int{}
	currentPermutation := make([]int, 0, len(numbers))
	isUsed := make([]bool, len(numbers))
	buildPermutations(numbers, currentPermutation, isUsed, &allPermutations)
	return allPermutations
}

// buildPermutations recursively constructs permutations by choosing unused elements
func buildPermutations(numbers []int, currentPermutation []int, isUsed []bool, allPermutations *[][]int) {
	if len(currentPermutation) == len(numbers) {
		completedPermutation := make([]int, len(currentPermutation))
		copy(completedPermutation, currentPermutation)
		*allPermutations = append(*allPermutations, completedPermutation)
		return
	}
	for i := 0; i < len(numbers); i++ {
		if isUsed[i] {
			continue
		}
		isUsed[i] = true
		currentPermutation = append(currentPermutation, numbers[i])
		buildPermutations(numbers, currentPermutation, isUsed, allPermutations)
		currentPermutation = currentPermutation[:len(currentPermutation)-1]
		isUsed[i] = false
	}
}
