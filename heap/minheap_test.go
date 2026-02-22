package heap

import "testing"

func TestNewMinHeap(t *testing.T) {
	intHeap := NewMinHeap[int]()

	if intHeap.Size() != 0 {
		t.Errorf("expected size 0, got %d", intHeap.Size())
	}
	if !intHeap.IsEmpty() {
		t.Error("expected new heap to be empty")
	}
}

func TestPushAndPeek(t *testing.T) {
	intHeap := NewMinHeap[int]()

	intHeap.Push(10)
	intHeap.Push(5)
	intHeap.Push(20)

	// The minimum element should be at the top
	topElement, found := intHeap.Peek()
	if !found {
		t.Error("expected peek to succeed")
	}
	if topElement != 5 {
		t.Errorf("expected minimum 5, got %d", topElement)
	}
	if intHeap.Size() != 3 {
		t.Errorf("expected size 3, got %d", intHeap.Size())
	}
}

func TestPopOrder(t *testing.T) {
	intHeap := NewMinHeap[int]()

	// Insert values in random order
	insertionValues := []int{15, 10, 20, 5, 30, 1, 25}
	for _, value := range insertionValues {
		intHeap.Push(value)
	}

	// Popping should yield values in ascending sorted order
	expectedOrder := []int{1, 5, 10, 15, 20, 25, 30}
	for orderIndex, expectedValue := range expectedOrder {
		poppedValue, found := intHeap.Pop()
		if !found {
			t.Errorf("expected pop at position %d to succeed", orderIndex)
		}
		if poppedValue != expectedValue {
			t.Errorf("expected %d at position %d, got %d", expectedValue, orderIndex, poppedValue)
		}
	}

	// Heap should be empty after popping all elements
	if !intHeap.IsEmpty() {
		t.Error("expected heap to be empty after popping all elements")
	}
}

func TestPopEmpty(t *testing.T) {
	intHeap := NewMinHeap[int]()

	_, found := intHeap.Pop()
	if found {
		t.Error("expected pop on empty heap to return false")
	}
}

func TestPeekEmpty(t *testing.T) {
	intHeap := NewMinHeap[int]()

	_, found := intHeap.Peek()
	if found {
		t.Error("expected peek on empty heap to return false")
	}
}

func TestDuplicateValues(t *testing.T) {
	intHeap := NewMinHeap[int]()

	intHeap.Push(5)
	intHeap.Push(5)
	intHeap.Push(3)
	intHeap.Push(3)

	expectedOrder := []int{3, 3, 5, 5}
	for orderIndex, expectedValue := range expectedOrder {
		poppedValue, found := intHeap.Pop()
		if !found {
			t.Errorf("expected pop at position %d to succeed", orderIndex)
		}
		if poppedValue != expectedValue {
			t.Errorf("expected %d at position %d, got %d", expectedValue, orderIndex, poppedValue)
		}
	}
}

func TestStringHeap(t *testing.T) {
	// Verify min-heap works with string type (lexicographic ordering)
	stringHeap := NewMinHeap[string]()

	stringHeap.Push("cherry")
	stringHeap.Push("apple")
	stringHeap.Push("banana")

	topElement, found := stringHeap.Peek()
	if !found {
		t.Error("expected peek to succeed")
	}
	if topElement != "apple" {
		t.Errorf("expected 'apple', got %s", topElement)
	}
}

func TestPushAfterPop(t *testing.T) {
	intHeap := NewMinHeap[int]()

	intHeap.Push(10)
	intHeap.Push(5)

	intHeap.Pop()

	// After popping 5, push 3 which should become the new minimum
	intHeap.Push(3)

	topElement, found := intHeap.Peek()
	if !found {
		t.Error("expected peek to succeed")
	}
	if topElement != 3 {
		t.Errorf("expected 3, got %d", topElement)
	}
}
