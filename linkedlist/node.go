package linkedlist

// ListNode represents a single node in a singly linked list
type ListNode struct {
	Value    int
	NextNode *ListNode
}

// NewListNode creates a new node with the given value
func NewListNode(value int) *ListNode {
	return &ListNode{Value: value, NextNode: nil}
}

// BuildLinkedList creates a linked list from a slice of values
func BuildLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	headNode := NewListNode(values[0])
	currentNode := headNode
	for i := 1; i < len(values); i++ {
		currentNode.NextNode = NewListNode(values[i])
		currentNode = currentNode.NextNode
	}
	return headNode
}

// ToSlice converts a linked list back to a slice for easy comparison
func ToSlice(headNode *ListNode) []int {
	result := []int{}
	currentNode := headNode
	for currentNode != nil {
		result = append(result, currentNode.Value)
		currentNode = currentNode.NextNode
	}
	return result
}
