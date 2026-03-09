package linkedlist

import "testing"

func TestReverseList(t *testing.T) {
	headNode := BuildLinkedList([]int{1, 2, 3, 4, 5})
	reversedHead := ReverseList(headNode)
	result := ToSlice(reversedHead)
	expected := []int{5, 4, 3, 2, 1}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("index %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}

func TestMergeTwoSortedLists(t *testing.T) {
	firstList := BuildLinkedList([]int{1, 3, 5})
	secondList := BuildLinkedList([]int{2, 4, 6})
	mergedHead := MergeTwoSortedLists(firstList, secondList)
	result := ToSlice(mergedHead)
	expected := []int{1, 2, 3, 4, 5, 6}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("index %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}

func TestHasCycle(t *testing.T) {
	headNode := BuildLinkedList([]int{1, 2, 3})
	if HasCycle(headNode) {
		t.Error("expected no cycle in linear list")
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	headNode := BuildLinkedList([]int{1, 2, 3, 4, 5})
	resultHead := RemoveNthFromEnd(headNode, 2)
	result := ToSlice(resultHead)
	if len(result) != 4 {
		t.Errorf("expected 4 elements, got %d", len(result))
	}
}

func TestIsPalindrome(t *testing.T) {
	headNode := BuildLinkedList([]int{1, 2, 2, 1})
	if !IsPalindrome(headNode) {
		t.Error("expected palindrome")
	}
}
