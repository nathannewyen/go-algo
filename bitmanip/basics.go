package bitmanip

// CountSetBits returns the number of 1-bits in an integer (Hamming weight)
func CountSetBits(number int) int {
	setBitCount := 0
	for number != 0 {
		// Clear the lowest set bit
		number = number & (number - 1)
		setBitCount++
	}
	return setBitCount
}

// IsPowerOfTwo checks if number is a power of two using bit manipulation
func IsPowerOfTwo(number int) bool {
	if number <= 0 {
		return false
	}
	// Power of 2 has exactly one set bit
	return number&(number-1) == 0
}

// SingleNumber finds the element that appears only once (others appear twice)
func SingleNumber(numbers []int) int {
	// XOR of all elements cancels out duplicates
	uniqueElement := 0
	for _, currentNumber := range numbers {
		uniqueElement ^= currentNumber
	}
	return uniqueElement
}
