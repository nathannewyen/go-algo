package twopointers

import "unicode"

// IsValidPalindrome checks if string is palindrome ignoring non-alphanumeric chars
func IsValidPalindrome(inputString string) bool {
	runeSlice := []rune(inputString)
	leftPointer := 0
	rightPointer := len(runeSlice) - 1

	for leftPointer < rightPointer {
		// Skip non-alphanumeric characters from left
		for leftPointer < rightPointer && !unicode.IsLetter(runeSlice[leftPointer]) && !unicode.IsDigit(runeSlice[leftPointer]) {
			leftPointer++
		}
		// Skip non-alphanumeric characters from right
		for leftPointer < rightPointer && !unicode.IsLetter(runeSlice[rightPointer]) && !unicode.IsDigit(runeSlice[rightPointer]) {
			rightPointer--
		}
		// Compare characters case-insensitively
		if unicode.ToLower(runeSlice[leftPointer]) != unicode.ToLower(runeSlice[rightPointer]) {
			return false
		}
		leftPointer++
		rightPointer--
	}
	return true
}
