package intervals

import "fmt"

// SummaryRanges groups consecutive numbers into ranges like "1->3", "5", "7->9"
func SummaryRanges(sortedNumbers []int) []string {
	ranges := []string{}
	if len(sortedNumbers) == 0 {
		return ranges
	}
	rangeStart := sortedNumbers[0]
	rangeEnd := sortedNumbers[0]
	for i := 1; i < len(sortedNumbers); i++ {
		// If current number continues the consecutive sequence
		if sortedNumbers[i] == rangeEnd+1 {
			rangeEnd = sortedNumbers[i]
		} else {
			// Finalize the current range and start a new one
			ranges = append(ranges, formatRange(rangeStart, rangeEnd))
			rangeStart = sortedNumbers[i]
			rangeEnd = sortedNumbers[i]
		}
	}
	ranges = append(ranges, formatRange(rangeStart, rangeEnd))
	return ranges
}

// formatRange creates the string representation of a range
func formatRange(start int, end int) string {
	if start == end {
		return fmt.Sprintf("%d", start)
	}
	return fmt.Sprintf("%d->%d", start, end)
}
