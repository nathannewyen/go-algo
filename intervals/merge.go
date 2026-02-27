package intervals

// MergeIntervals merges all overlapping intervals into non-overlapping intervals
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// Sort intervals by start time
	sortIntervals(intervals)
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		lastMerged := merged[len(merged)-1]
		currentInterval := intervals[i]
		// Check if current interval overlaps with last merged interval
		if currentInterval[0] <= lastMerged[1] {
			if currentInterval[1] > lastMerged[1] {
				lastMerged[1] = currentInterval[1]
			}
		} else {
			merged = append(merged, currentInterval)
		}
	}
	return merged
}

// sortIntervals sorts intervals by their start value using insertion sort
func sortIntervals(intervals [][]int) {
	for i := 1; i < len(intervals); i++ {
		key := intervals[i]
		j := i - 1
		for j >= 0 && intervals[j][0] > key[0] {
			intervals[j+1] = intervals[j]
			j--
		}
		intervals[j+1] = key
	}
}
