package stack

// MinStack supports push, pop, and retrieving the minimum in constant time
type MinStack struct {
	valueStack   []int
	minimumStack []int
}

// NewMinStack creates a new MinStack instance
func NewMinStack() *MinStack {
	return &MinStack{
		valueStack:   []int{},
		minimumStack: []int{},
	}
}

// Push adds a value and updates the minimum tracking stack
func (ms *MinStack) Push(value int) {
	ms.valueStack = append(ms.valueStack, value)
	// Track the running minimum at each level of the stack
	if len(ms.minimumStack) == 0 {
		ms.minimumStack = append(ms.minimumStack, value)
	} else {
		currentMinimum := ms.minimumStack[len(ms.minimumStack)-1]
		if value < currentMinimum {
			ms.minimumStack = append(ms.minimumStack, value)
		} else {
			ms.minimumStack = append(ms.minimumStack, currentMinimum)
		}
	}
}

// Pop removes the top element and its associated minimum
func (ms *MinStack) Pop() (int, bool) {
	if len(ms.valueStack) == 0 {
		return 0, false
	}
	topIndex := len(ms.valueStack) - 1
	topValue := ms.valueStack[topIndex]
	ms.valueStack = ms.valueStack[:topIndex]
	ms.minimumStack = ms.minimumStack[:topIndex]
	return topValue, true
}

// GetMin returns the current minimum value in the stack
func (ms *MinStack) GetMin() (int, bool) {
	if len(ms.minimumStack) == 0 {
		return 0, false
	}
	return ms.minimumStack[len(ms.minimumStack)-1], true
}
