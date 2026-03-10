package mathalgo

// FibonacciIterative returns the nth fibonacci number iteratively
func FibonacciIterative(position int) int {
	if position <= 1 {
		return position
	}
	previousValue := 0
	currentValue := 1
	for step := 2; step <= position; step++ {
		nextValue := previousValue + currentValue
		previousValue = currentValue
		currentValue = nextValue
	}
	return currentValue
}

// FibonacciMemoized returns the nth fibonacci number using memoization
func FibonacciMemoized(position int) int {
	memo := map[int]int{0: 0, 1: 1}
	return fibHelper(position, memo)
}

// fibHelper recursively computes fibonacci with memoization cache
func fibHelper(position int, memo map[int]int) int {
	if cachedValue, exists := memo[position]; exists {
		return cachedValue
	}
	computedValue := fibHelper(position-1, memo) + fibHelper(position-2, memo)
	memo[position] = computedValue
	return computedValue
}
