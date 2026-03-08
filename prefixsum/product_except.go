package prefixsum

// ProductExceptSelf returns array where each element is product of all others
func ProductExceptSelf(numbers []int) []int {
	arrayLength := len(numbers)
	productResult := make([]int, arrayLength)

	// First pass: build left-side running product for each position
	leftRunningProduct := 1
	for i := 0; i < arrayLength; i++ {
		productResult[i] = leftRunningProduct
		leftRunningProduct *= numbers[i]
	}

	// Second pass: multiply by right-side running product for each position
	rightRunningProduct := 1
	for i := arrayLength - 1; i >= 0; i-- {
		productResult[i] *= rightRunningProduct
		rightRunningProduct *= numbers[i]
	}
	return productResult
}
