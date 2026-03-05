package stringalgo

import (
	"fmt"
	"strconv"
	"strings"
)

// CompressString performs run-length encoding on a string
// Example: "aabcccccaaa" -> "a2b1c5a3"
func CompressString(inputString string) string {
	if len(inputString) == 0 {
		return inputString
	}

	var compressedBuilder strings.Builder
	consecutiveCount := 1

	for charIndex := 1; charIndex < len(inputString); charIndex++ {
		if inputString[charIndex] == inputString[charIndex-1] {
			consecutiveCount++
		} else {
			compressedBuilder.WriteByte(inputString[charIndex-1])
			compressedBuilder.WriteString(fmt.Sprintf("%d", consecutiveCount))
			consecutiveCount = 1
		}
	}
	// Write last character group
	compressedBuilder.WriteByte(inputString[len(inputString)-1])
	compressedBuilder.WriteString(fmt.Sprintf("%d", consecutiveCount))

	compressedResult := compressedBuilder.String()
	// Only return compressed if it is actually shorter
	if len(compressedResult) >= len(inputString) {
		return inputString
	}
	return compressedResult
}

// DecompressString reverses run-length encoding
func DecompressString(compressedString string) string {
	var decompressedBuilder strings.Builder
	charIndex := 0

	for charIndex < len(compressedString) {
		currentChar := compressedString[charIndex]
		charIndex++

		// Read the count digits
		countStart := charIndex
		for charIndex < len(compressedString) && compressedString[charIndex] >= 0 && compressedString[charIndex] <= 9 {
			charIndex++
		}
		repeatCount, _ := strconv.Atoi(compressedString[countStart:charIndex])

		for repeatIndex := 0; repeatIndex < repeatCount; repeatIndex++ {
			decompressedBuilder.WriteByte(currentChar)
		}
	}
	return decompressedBuilder.String()
}
