package sliding

// MaxSumSubarray finds the maximum sum of any contiguous subarray of the given size.
// Uses the sliding window technique for O(n) time complexity.
// Returns the maximum sum and true, or 0 and false if the window size exceeds the slice length.
func MaxSumSubarray(numbers []int, windowSize int) (int, bool) {
	if windowSize > len(numbers) || windowSize <= 0 {
		return 0, false
	}

	// Calculate the sum of the first window
	currentWindowSum := 0
	for initialIndex := 0; initialIndex < windowSize; initialIndex++ {
		currentWindowSum += numbers[initialIndex]
	}

	maxWindowSum := currentWindowSum

	// Slide the window: add the new element, remove the old element
	for slideIndex := windowSize; slideIndex < len(numbers); slideIndex++ {
		currentWindowSum += numbers[slideIndex] - numbers[slideIndex-windowSize]

		if currentWindowSum > maxWindowSum {
			maxWindowSum = currentWindowSum
		}
	}

	return maxWindowSum, true
}

// LongestUniqueSubstring finds the length of the longest substring
// without repeating characters using the sliding window technique.
func LongestUniqueSubstring(inputString string) int {
	if len(inputString) == 0 {
		return 0
	}

	// Track the last seen index of each character within the current window
	lastSeenIndex := make(map[byte]int)
	maxUniqueLength := 0
	windowStartIndex := 0

	for currentIndex := 0; currentIndex < len(inputString); currentIndex++ {
		currentCharacter := inputString[currentIndex]

		// If the character was seen before and is within the current window,
		// move the window start past the previous occurrence
		if previousIndex, alreadySeen := lastSeenIndex[currentCharacter]; alreadySeen && previousIndex >= windowStartIndex {
			windowStartIndex = previousIndex + 1
		}

		lastSeenIndex[currentCharacter] = currentIndex

		currentWindowLength := currentIndex - windowStartIndex + 1
		if currentWindowLength > maxUniqueLength {
			maxUniqueLength = currentWindowLength
		}
	}

	return maxUniqueLength
}

// MinWindowContaining finds the smallest substring of `sourceString` that
// contains all characters from `targetCharacters`.
// Returns the substring and true if found, empty string and false otherwise.
func MinWindowContaining(sourceString string, targetCharacters string) (string, bool) {
	if len(sourceString) == 0 || len(targetCharacters) == 0 {
		return "", false
	}

	// Count how many of each character we need from the target
	targetCharCounts := make(map[byte]int)
	for targetIndex := 0; targetIndex < len(targetCharacters); targetIndex++ {
		targetCharCounts[targetCharacters[targetIndex]]++
	}

	requiredCharCount := len(targetCharCounts)
	satisfiedCharCount := 0

	// Track character counts in the current window
	windowCharCounts := make(map[byte]int)
	windowStartIndex := 0
	minimumWindowLength := len(sourceString) + 1
	minimumWindowStartIndex := 0

	for expandIndex := 0; expandIndex < len(sourceString); expandIndex++ {
		expandedCharacter := sourceString[expandIndex]
		windowCharCounts[expandedCharacter]++

		// Check if adding this character satisfies a target character requirement
		if neededCount, isTarget := targetCharCounts[expandedCharacter]; isTarget {
			if windowCharCounts[expandedCharacter] == neededCount {
				satisfiedCharCount++
			}
		}

		// Try to shrink the window from the left while all requirements are met
		for satisfiedCharCount == requiredCharCount {
			currentWindowLength := expandIndex - windowStartIndex + 1
			if currentWindowLength < minimumWindowLength {
				minimumWindowLength = currentWindowLength
				minimumWindowStartIndex = windowStartIndex
			}

			shrunkCharacter := sourceString[windowStartIndex]
			windowCharCounts[shrunkCharacter]--

			if neededCount, isTarget := targetCharCounts[shrunkCharacter]; isTarget {
				if windowCharCounts[shrunkCharacter] < neededCount {
					satisfiedCharCount--
				}
			}

			windowStartIndex++
		}
	}

	if minimumWindowLength > len(sourceString) {
		return "", false
	}

	return sourceString[minimumWindowStartIndex : minimumWindowStartIndex+minimumWindowLength], true
}
