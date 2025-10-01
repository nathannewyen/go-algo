package backtracking

// PartitionIntoPalindromes finds all ways to partition string into palindromic substrings
func PartitionIntoPalindromes(inputString string) [][]string {
	allPartitions := [][]string{}
	currentPartition := []string{}
	findPalindromePartitions(inputString, 0, currentPartition, &allPartitions)
	return allPartitions
}

// findPalindromePartitions recursively partitions string starting from startIndex
func findPalindromePartitions(inputString string, startIndex int, currentPartition []string, allPartitions *[][]string) {
	if startIndex == len(inputString) {
		partitionCopy := make([]string, len(currentPartition))
		copy(partitionCopy, currentPartition)
		*allPartitions = append(*allPartitions, partitionCopy)
		return
	}

	for endIndex := startIndex + 1; endIndex <= len(inputString); endIndex++ {
		candidateSubstring := inputString[startIndex:endIndex]
		if isStringPalindrome(candidateSubstring) {
			currentPartition = append(currentPartition, candidateSubstring)
			findPalindromePartitions(inputString, endIndex, currentPartition, allPartitions)
			currentPartition = currentPartition[:len(currentPartition)-1]
		}
	}
}

// isStringPalindrome checks if a string reads the same forwards and backwards
func isStringPalindrome(candidate string) bool {
	leftIndex := 0
	rightIndex := len(candidate) - 1
	for leftIndex < rightIndex {
		if candidate[leftIndex] != candidate[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}
	return true
}
