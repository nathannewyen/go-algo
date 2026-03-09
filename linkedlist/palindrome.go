package linkedlist

// IsPalindrome checks if a linked list reads the same forwards and backwards
func IsPalindrome(headNode *ListNode) bool {
	if headNode == nil {
		return true
	}
	if headNode.NextNode == nil {
		return true
	}
	// Find the middle of the list using slow and fast pointers
	slowPointer := headNode
	fastPointer := headNode
	for fastPointer.NextNode != nil && fastPointer.NextNode.NextNode != nil {
		slowPointer = slowPointer.NextNode
		fastPointer = fastPointer.NextNode.NextNode
	}
	// Reverse the second half of the list starting after the middle
	secondHalfHead := ReverseList(slowPointer.NextNode)
	// Compare first half with reversed second half
	firstHalfPointer := headNode
	secondHalfPointer := secondHalfHead
	isPalindromeResult := true
	for secondHalfPointer != nil {
		if firstHalfPointer.Value != secondHalfPointer.Value {
			isPalindromeResult = false
			break
		}
		firstHalfPointer = firstHalfPointer.NextNode
		secondHalfPointer = secondHalfPointer.NextNode
	}
	// Restore the list by reversing the second half back
	slowPointer.NextNode = ReverseList(secondHalfHead)
	return isPalindromeResult
}
