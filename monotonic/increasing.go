package monotonic

// NextSmallerElement finds the next smaller element for each position
func NextSmallerElement(numbers []int) []int {
	elementCount := len(numbers)
	result := make([]int, elementCount)
	for i := range result {
		result[i] = -1
	}
	// Stack maintains indices in decreasing order of their values
	decreasingStack := []int{}
	for currentIndex := 0; currentIndex < elementCount; currentIndex++ {
		for len(decreasingStack) > 0 {
			topIndex := decreasingStack[len(decreasingStack)-1]
			if numbers[currentIndex] >= numbers[topIndex] {
				break
			}
			decreasingStack = decreasingStack[:len(decreasingStack)-1]
			result[topIndex] = numbers[currentIndex]
		}
		decreasingStack = append(decreasingStack, currentIndex)
	}
	return result
}

// PreviousSmallerElement finds the previous smaller element for each position
func PreviousSmallerElement(numbers []int) []int {
	elementCount := len(numbers)
	result := make([]int, elementCount)
	for i := range result {
		result[i] = -1
	}
	increasingStack := []int{}
	for currentIndex := 0; currentIndex < elementCount; currentIndex++ {
		for len(increasingStack) > 0 && numbers[increasingStack[len(increasingStack)-1]] >= numbers[currentIndex] {
			increasingStack = increasingStack[:len(increasingStack)-1]
		}
		if len(increasingStack) > 0 {
			result[currentIndex] = numbers[increasingStack[len(increasingStack)-1]]
		}
		increasingStack = append(increasingStack, currentIndex)
	}
	return result
}
