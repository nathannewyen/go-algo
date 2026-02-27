package intervals

// HasOverlap checks if any two intervals in the list overlap
func HasOverlap(intervals [][]int) bool {
	sortIntervals(intervals)
	// Compare each consecutive pair of intervals for overlap
	for i := 1; i < len(intervals); i++ {
		previousEnd := intervals[i-1][1]
		currentStart := intervals[i][0]
		if currentStart < previousEnd {
			return true
		}
	}
	return false
}

// CountOverlaps returns the number of overlapping interval pairs
func CountOverlaps(intervals [][]int) int {
	sortIntervals(intervals)
	overlapCount := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			overlapCount++
		}
	}
	return overlapCount
}
