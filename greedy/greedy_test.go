package greedy

import "testing"

func TestSelectMaxActivities(t *testing.T) {
	activityList := []Activity{
		{1, 4}, {3, 5}, {0, 6}, {5, 7}, {3, 9}, {5, 9}, {6, 10}, {8, 11},
	}
	selectedActivities := SelectMaxActivities(activityList)
	if len(selectedActivities) < 3 {
		t.Errorf("Expected at least 3 activities, got %d", len(selectedActivities))
	}
}

func TestCanReachEnd(t *testing.T) {
	if !CanReachEnd([]int{2, 3, 1, 1, 4}) {
		t.Error("Should be able to reach end")
	}
	if CanReachEnd([]int{3, 2, 1, 0, 4}) {
		t.Error("Should not be able to reach end")
	}
}

func TestMergeOverlappingIntervals(t *testing.T) {
	intervalList := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	mergedResult := MergeOverlappingIntervals(intervalList)
	if len(mergedResult) != 3 {
		t.Errorf("Expected 3 merged intervals, got %d", len(mergedResult))
	}
}
