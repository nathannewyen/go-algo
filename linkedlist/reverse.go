package linkedlist

// ReverseList reverses a singly linked list iteratively
func ReverseList(headNode *ListNode) *ListNode {
	var previousNode *ListNode
	currentNode := headNode
	for currentNode != nil {
		nextNode := currentNode.NextNode
		// Point current node backward to previous node
		currentNode.NextNode = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}

// ReverseListRecursive reverses a linked list using recursion
func ReverseListRecursive(headNode *ListNode) *ListNode {
	if headNode == nil {
		return nil
	}
	if headNode.NextNode == nil {
		return headNode
	}
	// Recursively reverse the rest of the list
	reversedHead := ReverseListRecursive(headNode.NextNode)
	// Make the next node point back to current node
	headNode.NextNode.NextNode = headNode
	headNode.NextNode = nil
	return reversedHead
}
