package stack

// Stack represents a LIFO data structure using a slice
type Stack struct {
	elements []int
}

// NewStack creates and returns an empty stack
func NewStack() *Stack {
	return &Stack{elements: []int{}}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	topIndex := len(s.elements) - 1
	topValue := s.elements[topIndex]
	s.elements = s.elements[:topIndex]
	return topValue, true
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty checks if the stack has no elements
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.elements)
}
