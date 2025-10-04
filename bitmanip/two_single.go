package bitmanip

// FindTwoSingleNumbers finds two elements that appear once when all others appear twice
func FindTwoSingleNumbers(numbers []int) [2]int {
	// XOR all numbers gives XOR of the two unique numbers
	combinedXOR := 0
	for _, currentNumber := range numbers {
		combinedXOR ^= currentNumber
	}

	// Find rightmost set bit to partition numbers into two groups
	rightmostSetBit := combinedXOR & (-combinedXOR)

	firstUniqueNumber := 0
	secondUniqueNumber := 0
	for _, currentNumber := range numbers {
		// Split into two groups based on the differentiating bit
		if currentNumber&rightmostSetBit != 0 {
			firstUniqueNumber ^= currentNumber
		} else {
			secondUniqueNumber ^= currentNumber
		}
	}
	return [2]int{firstUniqueNumber, secondUniqueNumber}
}
