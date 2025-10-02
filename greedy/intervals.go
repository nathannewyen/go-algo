package greedy

import "sort"

// Interval represents a range with start and end values
type Interval struct {
	Start int
	End   int
}

// MergeOverlappingIntervals combines intervals that overlap
func MergeOverlappingIntervals(intervalList []Interval) []Interval {
	if len(intervalList) <= 1 {
		return intervalList
	}

	// Sort intervals by start time
	sort.Slice(intervalList, func(i int, j int) bool {
		return intervalList[i].Start < intervalList[j].Start
	})

	mergedIntervals := []Interval{intervalList[0]}

	for intervalIndex := 1; intervalIndex < len(intervalList); intervalIndex++ {
		lastMergedInterval := &mergedIntervals[len(mergedIntervals)-1]
		currentInterval := intervalList[intervalIndex]

		// If current interval overlaps with last merged, extend it
		if currentInterval.Start <= lastMergedInterval.End {
			if currentInterval.End > lastMergedInterval.End {
				lastMergedInterval.End = currentInterval.End
			}
		} else {
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}
	return mergedIntervals
}
