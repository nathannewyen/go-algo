package recursion

// LetterCombinations returns all possible letter combos for a phone number string
func LetterCombinations(phoneDigits string) []string {
	if len(phoneDigits) == 0 {
		return []string{}
	}
	// Map each digit to its corresponding letters on a phone keypad
	digitToLetters := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	allCombinations := []string{}
	buildLetterCombinations(phoneDigits, 0, "", digitToLetters, &allCombinations)
	return allCombinations
}

// buildLetterCombinations recursively appends letters for each digit
func buildLetterCombinations(phoneDigits string, digitIndex int, currentCombo string, digitToLetters map[byte]string, allCombinations *[]string) {
	if digitIndex == len(phoneDigits) {
		*allCombinations = append(*allCombinations, currentCombo)
		return
	}
	currentDigit := phoneDigits[digitIndex]
	availableLetters := digitToLetters[currentDigit]
	for i := 0; i < len(availableLetters); i++ {
		chosenLetter := string(availableLetters[i])
		buildLetterCombinations(phoneDigits, digitIndex+1, currentCombo+chosenLetter, digitToLetters, allCombinations)
	}
}
