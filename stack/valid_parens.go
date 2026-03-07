package stack

// IsValidParentheses checks if a string of brackets is properly balanced
func IsValidParentheses(expression string) bool {
	bracketStack := NewStack()
	// Map closing brackets to their opening counterparts
	matchingBrackets := map[byte]byte{')': '(', '}': '{', ']': '['}
	openingBrackets := map[byte]bool{'(': true, '{': true, '[': true}

	for i := 0; i < len(expression); i++ {
		currentChar := expression[i]
		if openingBrackets[currentChar] {
			bracketStack.Push(int(currentChar))
		} else if expectedOpening, isClosing := matchingBrackets[currentChar]; isClosing {
			topValue, hasValue := bracketStack.Pop()
			if !hasValue {
				return false
			}
			if byte(topValue) != expectedOpening {
				return false
			}
		}
	}
	return bracketStack.IsEmpty()
}
