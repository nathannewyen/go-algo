package heap

import "golang.org/x/exp/constraints"

// MinHeap implements a min-heap where the smallest element is always at the root.
// Uses a slice-based binary heap with index math for parent/child relationships.
type MinHeap[T constraints.Ordered] struct {
	elements []T
}

// NewMinHeap creates and returns an empty min-heap.
func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{
		elements: make([]T, 0),
	}
}

// parentIndex returns the index of the parent for a given child index.
func parentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// leftChildIndex returns the index of the left child for a given parent index.
func leftChildIndex(parentIdx int) int {
	return 2*parentIdx + 1
}

// rightChildIndex returns the index of the right child for a given parent index.
func rightChildIndex(parentIdx int) int {
	return 2*parentIdx + 2
}

// Push adds an element to the heap and restores the heap property by sifting up.
func (h *MinHeap[T]) Push(value T) {
	h.elements = append(h.elements, value)
	h.siftUp(len(h.elements) - 1)
}

// Pop removes and returns the minimum element from the heap.
// Returns the zero value and false if the heap is empty.
func (h *MinHeap[T]) Pop() (T, bool) {
	if len(h.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	minimumElement := h.elements[0]
	lastIndex := len(h.elements) - 1

	// Move the last element to the root and shrink the slice
	h.elements[0] = h.elements[lastIndex]
	h.elements = h.elements[:lastIndex]

	// Restore heap property by sifting the new root down
	if len(h.elements) > 0 {
		h.siftDown(0)
	}

	return minimumElement, true
}

// Peek returns the minimum element without removing it.
// Returns the zero value and false if the heap is empty.
func (h *MinHeap[T]) Peek() (T, bool) {
	if len(h.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return h.elements[0], true
}

// Size returns the number of elements in the heap.
func (h *MinHeap[T]) Size() int {
	return len(h.elements)
}

// IsEmpty returns true if the heap has no elements.
func (h *MinHeap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

// siftUp moves an element up the heap until the min-heap property is restored.
// Called after inserting a new element at the bottom.
func (h *MinHeap[T]) siftUp(currentIndex int) {
	for currentIndex > 0 {
		parentIdx := parentIndex(currentIndex)

		// If the current element is not smaller than its parent, heap property holds
		if h.elements[currentIndex] >= h.elements[parentIdx] {
			break
		}

		// Swap the current element with its parent
		h.elements[currentIndex], h.elements[parentIdx] = h.elements[parentIdx], h.elements[currentIndex]
		currentIndex = parentIdx
	}
}

// siftDown moves an element down the heap until the min-heap property is restored.
// Called after removing the root and replacing it with the last element.
func (h *MinHeap[T]) siftDown(currentIndex int) {
	totalElements := len(h.elements)

	for {
		smallestIndex := currentIndex
		leftIdx := leftChildIndex(currentIndex)
		rightIdx := rightChildIndex(currentIndex)

		// Check if the left child is smaller than the current smallest
		if leftIdx < totalElements && h.elements[leftIdx] < h.elements[smallestIndex] {
			smallestIndex = leftIdx
		}

		// Check if the right child is smaller than the current smallest
		if rightIdx < totalElements && h.elements[rightIdx] < h.elements[smallestIndex] {
			smallestIndex = rightIdx
		}

		// If the current element is already the smallest, heap property holds
		if smallestIndex == currentIndex {
			break
		}

		h.elements[currentIndex], h.elements[smallestIndex] = h.elements[smallestIndex], h.elements[currentIndex]
		currentIndex = smallestIndex
	}
}
