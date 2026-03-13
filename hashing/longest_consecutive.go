package hashing

// LongestConsecutiveSequence finds the length of the longest consecutive element sequence
func LongestConsecutiveSequence(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	// Build a set for O(1) existence checks
	numberSet := map[int]bool{}
	for _, num := range numbers {
		numberSet[num] = true
	}
	longestStreak := 0
	for currentNumber := range numberSet {
		// Only start counting from the beginning of a sequence
		isSequenceStart := !numberSet[currentNumber-1]
		if isSequenceStart {
			sequenceLength := 1
			nextNumber := currentNumber + 1
			for numberSet[nextNumber] {
				sequenceLength++
				nextNumber++
			}
			if sequenceLength > longestStreak {
				longestStreak = sequenceLength
			}
		}
	}
	return longestStreak
}
