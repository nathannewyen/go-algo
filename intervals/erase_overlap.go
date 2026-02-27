package intervals

// EraseOverlapIntervals returns minimum intervals to remove for non-overlapping set
func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	// Sort by end time for greedy interval scheduling
	for i := 1; i < len(intervals); i++ {
		key := intervals[i]
		j := i - 1
		for j >= 0 && intervals[j][1] > key[1] {
			intervals[j+1] = intervals[j]
			j--
		}
		intervals[j+1] = key
	}
	removedCount := 0
	previousEnd := intervals[0][1]
	// Greedily keep intervals that start after the previous end
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < previousEnd {
			removedCount++
		} else {
			previousEnd = intervals[i][1]
		}
	}
	return removedCount
}
