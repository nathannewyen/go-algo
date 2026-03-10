package mathalgo

// Factorial computes n! iteratively
func Factorial(number int) int {
	if number <= 1 {
		return 1
	}
	result := 1
	for multiplier := 2; multiplier <= number; multiplier++ {
		result *= multiplier
	}
	return result
}

// Combinations computes C(n, k) = n! / (k! * (n-k)!)
func Combinations(totalItems int, chooseCount int) int {
	if chooseCount > totalItems {
		return 0
	}
	if chooseCount == 0 {
		return 1
	}
	// Use the smaller of k and n-k for efficiency
	if chooseCount > totalItems-chooseCount {
		chooseCount = totalItems - chooseCount
	}
	result := 1
	for i := 0; i < chooseCount; i++ {
		result *= (totalItems - i)
		result /= (i + 1)
	}
	return result
}
