package hashing

// TwoSum finds indices of two numbers that add up to the target
func TwoSum(numbers []int, targetSum int) []int {
	// Map each seen value to its index for O(1) complement lookup
	valueToIndex := map[int]int{}
	for currentIndex, currentValue := range numbers {
		complementValue := targetSum - currentValue
		if complementIndex, exists := valueToIndex[complementValue]; exists {
			return []int{complementIndex, currentIndex}
		}
		valueToIndex[currentValue] = currentIndex
	}
	return []int{}
}
