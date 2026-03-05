package stringalgo

const hashPrime = 101

// RabinKarpSearch finds pattern occurrences using rolling hash technique
func RabinKarpSearch(textContent string, searchPattern string) []int {
	matchPositions := []int{}
	patternLength := len(searchPattern)
	textLength := len(textContent)

	if patternLength > textLength {
		return matchPositions
	}

	// Calculate hash for pattern and first window of text
	patternHash := computeRollingHash(searchPattern, patternLength)
	windowHash := computeRollingHash(textContent, patternLength)

	// Precompute highest power of base for rolling hash removal
	highestPower := 1
	for powerIndex := 0; powerIndex < patternLength-1; powerIndex++ {
		highestPower = (highestPower * 256) % hashPrime
	}

	for windowStart := 0; windowStart <= textLength-patternLength; windowStart++ {
		if patternHash == windowHash {
			// Hash match - verify character by character
			if textContent[windowStart:windowStart+patternLength] == searchPattern {
				matchPositions = append(matchPositions, windowStart)
			}
		}
		// Roll the hash window forward
		if windowStart < textLength-patternLength {
			windowHash = rollHashForward(windowHash, textContent, windowStart, patternLength, highestPower)
		}
	}
	return matchPositions
}

// computeRollingHash calculates initial hash value for a string
func computeRollingHash(content string, length int) int {
	hashValue := 0
	for charIndex := 0; charIndex < length; charIndex++ {
		hashValue = (hashValue*256 + int(content[charIndex])) % hashPrime
	}
	return hashValue
}

// rollHashForward updates hash by removing leftmost char and adding new rightmost char
func rollHashForward(currentHash int, textContent string, windowStart int, patternLength int, highestPower int) int {
	newHash := currentHash
	newHash = (newHash - int(textContent[windowStart])*highestPower) % hashPrime
	newHash = (newHash*256 + int(textContent[windowStart+patternLength])) % hashPrime
	if newHash < 0 {
		newHash += hashPrime
	}
	return newHash
}
