package bitmanip

// GetBitAtPosition returns the bit value at the given position
func GetBitAtPosition(number int, bitPosition int) int {
	return (number >> bitPosition) & 1
}

// SetBitAtPosition sets the bit at the given position to 1
func SetBitAtPosition(number int, bitPosition int) int {
	return number | (1 << bitPosition)
}

// ClearBitAtPosition sets the bit at the given position to 0
func ClearBitAtPosition(number int, bitPosition int) int {
	bitMask := ^(1 << bitPosition)
	return number & bitMask
}

// ToggleBitAtPosition flips the bit at the given position
func ToggleBitAtPosition(number int, bitPosition int) int {
	return number ^ (1 << bitPosition)
}

// HammingDistance counts differing bit positions between two numbers
func HammingDistance(numberA int, numberB int) int {
	// XOR gives 1 where bits differ
	differingBits := numberA ^ numberB
	return CountSetBits(differingBits)
}
