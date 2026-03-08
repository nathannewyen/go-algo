package prefixsum

// BuildPrefixSum creates a prefix sum array for efficient range sum queries
func BuildPrefixSum(numbers []int) []int {
	prefixSumArray := make([]int, len(numbers)+1)
	// Each position stores the cumulative sum up to that index
	for i := 0; i < len(numbers); i++ {
		prefixSumArray[i+1] = prefixSumArray[i] + numbers[i]
	}
	return prefixSumArray
}

// RangeSum returns the sum of elements between leftIndex and rightIndex inclusive
func RangeSum(prefixSumArray []int, leftIndex int, rightIndex int) int {
	return prefixSumArray[rightIndex+1] - prefixSumArray[leftIndex]
}
