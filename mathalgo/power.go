package mathalgo

// FastPower computes base raised to exponent using binary exponentiation
func FastPower(baseValue int, exponentValue int) int {
	result := 1
	currentBase := baseValue
	remainingExponent := exponentValue
	// Square the base and halve the exponent at each step
	for remainingExponent > 0 {
		if remainingExponent%2 == 1 {
			result *= currentBase
		}
		currentBase *= currentBase
		remainingExponent /= 2
	}
	return result
}

// FastPowerMod computes (base^exp) % modulus efficiently
func FastPowerMod(baseValue int, exponentValue int, modulusValue int) int {
	result := 1
	currentBase := baseValue % modulusValue
	remainingExponent := exponentValue
	for remainingExponent > 0 {
		if remainingExponent%2 == 1 {
			result = (result * currentBase) % modulusValue
		}
		currentBase = (currentBase * currentBase) % modulusValue
		remainingExponent /= 2
	}
	return result
}
