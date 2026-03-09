package linkedlist

// FindMiddleNode returns the middle node of a linked list
func FindMiddleNode(headNode *ListNode) *ListNode {
	slowPointer := headNode
	fastPointer := headNode
	// When fast pointer reaches end, slow pointer is at middle
	for fastPointer != nil && fastPointer.NextNode != nil {
		slowPointer = slowPointer.NextNode
		fastPointer = fastPointer.NextNode.NextNode
	}
	return slowPointer
}
