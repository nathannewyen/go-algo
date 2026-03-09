package linkedlist

// HasCycle detects if a linked list contains a cycle using Floyd tortoise and hare
func HasCycle(headNode *ListNode) bool {
	if headNode == nil {
		return false
	}
	slowPointer := headNode
	fastPointer := headNode
	// Fast pointer moves two steps while slow moves one
	for fastPointer != nil && fastPointer.NextNode != nil {
		slowPointer = slowPointer.NextNode
		fastPointer = fastPointer.NextNode.NextNode
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}

// FindCycleStart returns the node where the cycle begins, or nil if no cycle
func FindCycleStart(headNode *ListNode) *ListNode {
	slowPointer := headNode
	fastPointer := headNode
	// Phase 1: detect if cycle exists
	hasCycle := false
	for fastPointer != nil && fastPointer.NextNode != nil {
		slowPointer = slowPointer.NextNode
		fastPointer = fastPointer.NextNode.NextNode
		if slowPointer == fastPointer {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	// Phase 2: find cycle entry point by resetting one pointer to head
	entryPointer := headNode
	for entryPointer != slowPointer {
		entryPointer = entryPointer.NextNode
		slowPointer = slowPointer.NextNode
	}
	return entryPointer
}
