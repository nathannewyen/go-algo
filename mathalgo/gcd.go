package mathalgo

// GreatestCommonDivisor computes the GCD of two numbers using Euclidean algorithm
func GreatestCommonDivisor(firstNumber int, secondNumber int) int {
	for secondNumber != 0 {
		remainder := firstNumber % secondNumber
		firstNumber = secondNumber
		secondNumber = remainder
	}
	return firstNumber
}

// LeastCommonMultiple computes the LCM using the GCD relationship
func LeastCommonMultiple(firstNumber int, secondNumber int) int {
	gcdValue := GreatestCommonDivisor(firstNumber, secondNumber)
	return (firstNumber / gcdValue) * secondNumber
}
