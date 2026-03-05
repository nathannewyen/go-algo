package stringalgo

// LongestPalindromicSubstring finds the longest palindrome within a string
// Uses expand-around-center approach: O(n^2) time, O(1) space
func LongestPalindromicSubstring(inputString string) string {
	if len(inputString) < 2 {
		return inputString
	}

	longestStart := 0
	longestLength := 1

	for centerIndex := 0; centerIndex < len(inputString); centerIndex++ {
		// Check odd-length palindromes centered at centerIndex
		oddStart, oddLength := expandFromCenter(inputString, centerIndex, centerIndex)
		if oddLength > longestLength {
			longestStart = oddStart
			longestLength = oddLength
		}

		// Check even-length palindromes centered between centerIndex and centerIndex+1
		evenStart, evenLength := expandFromCenter(inputString, centerIndex, centerIndex+1)
		if evenLength > longestLength {
			longestStart = evenStart
			longestLength = evenLength
		}
	}
	return inputString[longestStart : longestStart+longestLength]
}

// expandFromCenter expands outward from center while characters match
func expandFromCenter(inputString string, leftBound int, rightBound int) (int, int) {
	for leftBound >= 0 && rightBound < len(inputString) && inputString[leftBound] == inputString[rightBound] {
		leftBound--
		rightBound++
	}
	// Return start position and length of palindrome found
	return leftBound + 1, rightBound - leftBound - 1
}
