package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, ok := s.Pop()
	if !ok {
		t.Fatal("expected pop to succeed")
	}
	if val != 3 {
		t.Errorf("expected 3, got %d", val)
	}
}

func TestValidParentheses(t *testing.T) {
	if !IsValidParentheses("()[]{}") {
		t.Error("expected valid parentheses")
	}
	if IsValidParentheses("(]") {
		t.Error("expected invalid parentheses")
	}
}

func TestMinStack(t *testing.T) {
	ms := NewMinStack()
	ms.Push(3)
	ms.Push(1)
	ms.Push(2)
	minVal, _ := ms.GetMin()
	if minVal != 1 {
		t.Errorf("expected min 1, got %d", minVal)
	}
}

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	result := EvalRPN(tokens)
	if result != 9 {
		t.Errorf("expected 9, got %d", result)
	}
}

func TestDailyTemperatures(t *testing.T) {
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := DailyTemperatures(temps)
	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("index %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}
