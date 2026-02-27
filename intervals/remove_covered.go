package intervals

// RemoveCoveredIntervals returns the count of intervals that are not covered by another
func RemoveCoveredIntervals(intervals [][]int) int {
	// Sort by start ascending, then by end descending for same start
	for i := 1; i < len(intervals); i++ {
		key := intervals[i]
		j := i - 1
		for j >= 0 && (intervals[j][0] > key[0] || (intervals[j][0] == key[0] && intervals[j][1] < key[1])) {
			intervals[j+1] = intervals[j]
			j--
		}
		intervals[j+1] = key
	}
	remainingCount := 1
	largestEnd := intervals[0][1]
	// An interval is covered if its end is within the largest end seen so far
	for i := 1; i < len(intervals); i++ {
		currentEnd := intervals[i][1]
		if currentEnd > largestEnd {
			remainingCount++
			largestEnd = currentEnd
		}
	}
	return remainingCount
}
