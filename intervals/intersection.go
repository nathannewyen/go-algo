package intervals

// IntervalIntersection finds the intersection of two sorted interval lists
func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	intersections := [][]int{}
	firstPointer := 0
	secondPointer := 0
	for firstPointer < len(firstList) && secondPointer < len(secondList) {
		// Find the overlap boundaries between current intervals
		overlapStart := firstList[firstPointer][0]
		if secondList[secondPointer][0] > overlapStart {
			overlapStart = secondList[secondPointer][0]
		}
		overlapEnd := firstList[firstPointer][1]
		if secondList[secondPointer][1] < overlapEnd {
			overlapEnd = secondList[secondPointer][1]
		}
		// If valid overlap exists, add to result
		if overlapStart <= overlapEnd {
			intersections = append(intersections, []int{overlapStart, overlapEnd})
		}
		// Advance the pointer with the smaller end value
		if firstList[firstPointer][1] < secondList[secondPointer][1] {
			firstPointer++
		} else {
			secondPointer++
		}
	}
	return intersections
}
