package bitmanip

// ReverseBits reverses the bit order of a 32-bit unsigned integer
func ReverseBits(number uint32) uint32 {
	var reversedResult uint32
	for bitIndex := 0; bitIndex < 32; bitIndex++ {
		// Shift result left and add lowest bit of number
		reversedResult = (reversedResult << 1) | (number & 1)
		number >>= 1
	}
	return reversedResult
}

// FindMissingNumber finds the missing number in range [0, n] using XOR
func FindMissingNumber(numbers []int) int {
	expectedLength := len(numbers)
	xorResult := expectedLength

	for currentIndex := 0; currentIndex < expectedLength; currentIndex++ {
		// XOR with index and value to cancel out pairs
		xorResult ^= currentIndex ^ numbers[currentIndex]
	}
	return xorResult
}

// CountBitsInRange counts total set bits for all numbers from 0 to n
func CountBitsInRange(upperBound int) int {
	totalSetBits := 0
	for currentNumber := 0; currentNumber <= upperBound; currentNumber++ {
		totalSetBits += CountSetBits(currentNumber)
	}
	return totalSetBits
}
