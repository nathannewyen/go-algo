package linkedlist

// MergeTwoSortedLists merges two sorted linked lists into one sorted list
func MergeTwoSortedLists(firstList *ListNode, secondList *ListNode) *ListNode {
	// Sentinel node simplifies edge cases for building the merged list
	sentinelNode := &ListNode{}
	mergedTail := sentinelNode

	firstPointer := firstList
	secondPointer := secondList

	for firstPointer != nil && secondPointer != nil {
		if firstPointer.Value <= secondPointer.Value {
			mergedTail.NextNode = firstPointer
			firstPointer = firstPointer.NextNode
		} else {
			mergedTail.NextNode = secondPointer
			secondPointer = secondPointer.NextNode
		}
		mergedTail = mergedTail.NextNode
	}
	// Attach remaining nodes from whichever list is not exhausted
	if firstPointer != nil {
		mergedTail.NextNode = firstPointer
	} else {
		mergedTail.NextNode = secondPointer
	}
	return sentinelNode.NextNode
}
