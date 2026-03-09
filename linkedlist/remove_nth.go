package linkedlist

// RemoveNthFromEnd removes the nth node from the end of the linked list
func RemoveNthFromEnd(headNode *ListNode, positionFromEnd int) *ListNode {
	// Sentinel node handles edge case of removing the head
	sentinelNode := &ListNode{NextNode: headNode}
	leadPointer := sentinelNode
	trailPointer := sentinelNode

	// Advance lead pointer by n+1 positions to create the gap
	for step := 0; step <= positionFromEnd; step++ {
		leadPointer = leadPointer.NextNode
	}
	// Move both pointers until lead reaches the end
	for leadPointer != nil {
		leadPointer = leadPointer.NextNode
		trailPointer = trailPointer.NextNode
	}
	// Trail pointer is now just before the node to remove
	trailPointer.NextNode = trailPointer.NextNode.NextNode
	return sentinelNode.NextNode
}
