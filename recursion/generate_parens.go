package recursion

// GenerateParentheses generates all valid combinations of n pairs of parentheses
func GenerateParentheses(pairCount int) []string {
	validCombinations := []string{}
	buildParentheses(pairCount, 0, 0, "", &validCombinations)
	return validCombinations
}

// buildParentheses recursively builds valid parentheses strings
func buildParentheses(pairCount int, openCount int, closeCount int, currentString string, validCombinations *[]string) {
	if len(currentString) == pairCount*2 {
		*validCombinations = append(*validCombinations, currentString)
		return
	}
	// Can add open paren if we haven't used all available
	if openCount < pairCount {
		buildParentheses(pairCount, openCount+1, closeCount, currentString+"(", validCombinations)
	}
	// Can add close paren only if it won't exceed open count
	if closeCount < openCount {
		buildParentheses(pairCount, openCount, closeCount+1, currentString+")", validCombinations)
	}
}
