package greedy

import "sort"

// Activity represents a task with a start and finish time
type Activity struct {
	StartTime  int
	FinishTime int
}

// SelectMaxActivities picks the maximum number of non-overlapping activities
func SelectMaxActivities(activityList []Activity) []Activity {
	// Sort activities by finish time (greedy choice)
	sort.Slice(activityList, func(i int, j int) bool {
		return activityList[i].FinishTime < activityList[j].FinishTime
	})

	selectedActivities := []Activity{activityList[0]}
	lastFinishTime := activityList[0].FinishTime

	for activityIndex := 1; activityIndex < len(activityList); activityIndex++ {
		// Select activity if it starts after the last selected one finishes
		if activityList[activityIndex].StartTime >= lastFinishTime {
			selectedActivities = append(selectedActivities, activityList[activityIndex])
			lastFinishTime = activityList[activityIndex].FinishTime
		}
	}
	return selectedActivities
}
