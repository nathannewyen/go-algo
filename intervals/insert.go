package intervals

// InsertInterval inserts a new interval into a sorted list of non-overlapping intervals
func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	insertIndex := 0
	// Add all intervals that come before the new interval
	for insertIndex < len(intervals) && intervals[insertIndex][1] < newInterval[0] {
		result = append(result, intervals[insertIndex])
		insertIndex++
	}
	// Merge overlapping intervals with the new interval
	for insertIndex < len(intervals) && intervals[insertIndex][0] <= newInterval[1] {
		if intervals[insertIndex][0] < newInterval[0] {
			newInterval[0] = intervals[insertIndex][0]
		}
		if intervals[insertIndex][1] > newInterval[1] {
			newInterval[1] = intervals[insertIndex][1]
		}
		insertIndex++
	}
	result = append(result, newInterval)
	// Add remaining intervals after the merged region
	for insertIndex < len(intervals) {
		result = append(result, intervals[insertIndex])
		insertIndex++
	}
	return result
}
