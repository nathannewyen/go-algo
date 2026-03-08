package prefixsum

// MaxContiguousSum finds the maximum sum of any contiguous subarray using Kadane algorithm
func MaxContiguousSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	currentSubarraySum := numbers[0]
	globalMaximumSum := numbers[0]

	for i := 1; i < len(numbers); i++ {
		// Either extend the current subarray or start fresh from current element
		if currentSubarraySum+numbers[i] > numbers[i] {
			currentSubarraySum = currentSubarraySum + numbers[i]
		} else {
			currentSubarraySum = numbers[i]
		}
		if currentSubarraySum > globalMaximumSum {
			globalMaximumSum = currentSubarraySum
		}
	}
	return globalMaximumSum
}
