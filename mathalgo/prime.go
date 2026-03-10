package mathalgo

// IsPrime checks if a number is prime using trial division
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}
	if number < 4 {
		return true
	}
	if number%2 == 0 {
		return false
	}
	// Only need to check odd divisors up to square root
	divisor := 3
	for divisor*divisor <= number {
		if number%divisor == 0 {
			return false
		}
		divisor += 2
	}
	return true
}

// SieveOfEratosthenes returns all primes up to the given limit
func SieveOfEratosthenes(upperLimit int) []int {
	if upperLimit < 2 {
		return []int{}
	}
	// Boolean array where true means the index is composite (not prime)
	isComposite := make([]bool, upperLimit+1)
	isComposite[0] = true
	isComposite[1] = true
	for candidate := 2; candidate*candidate <= upperLimit; candidate++ {
		if !isComposite[candidate] {
			// Mark all multiples of this prime as composite
			for multiple := candidate * candidate; multiple <= upperLimit; multiple += candidate {
				isComposite[multiple] = true
			}
		}
	}
	primeNumbers := []int{}
	for number := 2; number <= upperLimit; number++ {
		if !isComposite[number] {
			primeNumbers = append(primeNumbers, number)
		}
	}
	return primeNumbers
}
