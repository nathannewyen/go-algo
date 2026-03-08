package prefixsum

// SubarraySumEqualsK counts subarrays whose elements sum to the target value
func SubarraySumEqualsK(numbers []int, targetSum int) int {
	// Track how many times each prefix sum has occurred
	prefixSumFrequency := map[int]int{0: 1}
	runningSum := 0
	subarrayCount := 0

	for _, currentNumber := range numbers {
		runningSum += currentNumber
		// If (runningSum - targetSum) was seen before, those positions form valid subarrays
		neededPrefixSum := runningSum - targetSum
		if frequency, exists := prefixSumFrequency[neededPrefixSum]; exists {
			subarrayCount += frequency
		}
		prefixSumFrequency[runningSum]++
	}
	return subarrayCount
}
