package stack

// EvalRPN evaluates an arithmetic expression in reverse polish notation
func EvalRPN(tokens []string) int {
	operandStack := []int{}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			// Pop two operands for the binary operation
			rightOperand := operandStack[len(operandStack)-1]
			leftOperand := operandStack[len(operandStack)-2]
			operandStack = operandStack[:len(operandStack)-2]
			var operationResult int
			switch token {
			case "+":
				operationResult = leftOperand + rightOperand
			case "-":
				operationResult = leftOperand - rightOperand
			case "*":
				operationResult = leftOperand * rightOperand
			case "/":
				operationResult = leftOperand / rightOperand
			}
			operandStack = append(operandStack, operationResult)
		default:
			// Parse number token and push onto operand stack
			numberValue := parseInteger(token)
			operandStack = append(operandStack, numberValue)
		}
	}
	return operandStack[0]
}

// parseInteger converts a string to an integer value
func parseInteger(numberString string) int {
	isNegative := false
	startIndex := 0
	if numberString[0] == '-' {
		isNegative = true
		startIndex = 1
	}
	result := 0
	for i := startIndex; i < len(numberString); i++ {
		digitValue := int(numberString[i] - '0')
		result = result*10 + digitValue
	}
	if isNegative {
		return -result
	}
	return result
}
