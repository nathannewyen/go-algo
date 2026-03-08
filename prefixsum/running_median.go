package prefixsum

// RunningSum returns a new array where each element is the cumulative sum
func RunningSum(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	cumulativeSums := make([]int, len(numbers))
	cumulativeSums[0] = numbers[0]
	for i := 1; i < len(numbers); i++ {
		cumulativeSums[i] = cumulativeSums[i-1] + numbers[i]
	}
	return cumulativeSums
}

// PivotIndex finds the index where left sum equals right sum
func PivotIndex(numbers []int) int {
	totalSum := 0
	for _, num := range numbers {
		totalSum += num
	}
	leftSideSum := 0
	for i, num := range numbers {
		rightSideSum := totalSum - leftSideSum - num
		if leftSideSum == rightSideSum {
			return i
		}
		leftSideSum += num
	}
	return -1
}
