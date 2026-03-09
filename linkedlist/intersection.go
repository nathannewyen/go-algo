package linkedlist

// FindIntersectionNode returns the node where two linked lists first intersect
func FindIntersectionNode(headNodeA *ListNode, headNodeB *ListNode) *ListNode {
	if headNodeA == nil {
		return nil
	}
	if headNodeB == nil {
		return nil
	}
	pointerA := headNodeA
	pointerB := headNodeB
	// When a pointer reaches the end, redirect it to the other list head
	// Both pointers will traverse the same total distance and meet at intersection
	for pointerA != pointerB {
		if pointerA == nil {
			pointerA = headNodeB
		} else {
			pointerA = pointerA.NextNode
		}
		if pointerB == nil {
			pointerB = headNodeA
		} else {
			pointerB = pointerB.NextNode
		}
	}
	return pointerA
}
