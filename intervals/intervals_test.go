package intervals

import "testing"

func TestMergeIntervals(t *testing.T) {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := MergeIntervals(input)
	if len(result) != 3 {
		t.Errorf("expected 3 merged intervals, got %d", len(result))
	}
}

func TestInsertInterval(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	result := InsertInterval(intervals, newInterval)
	if len(result) != 2 {
		t.Errorf("expected 2 intervals after insert, got %d", len(result))
	}
}

func TestCanAttendAllMeetings(t *testing.T) {
	meetings := [][]int{{0, 30}, {5, 10}, {15, 20}}
	if CanAttendAllMeetings(meetings) {
		t.Error("expected false for overlapping meetings")
	}
}

func TestMinMeetingRooms(t *testing.T) {
	meetings := [][]int{{0, 30}, {5, 10}, {15, 20}}
	roomsNeeded := MinMeetingRooms(meetings)
	if roomsNeeded != 2 {
		t.Errorf("expected 2 rooms, got %d", roomsNeeded)
	}
}

func TestRemoveCoveredIntervals(t *testing.T) {
	intervals := [][]int{{1, 4}, {3, 6}, {2, 8}}
	remainingCount := RemoveCoveredIntervals(intervals)
	if remainingCount != 2 {
		t.Errorf("expected 2 remaining, got %d", remainingCount)
	}
}

func TestIntervalIntersection(t *testing.T) {
	firstList := [][]int{{0, 2}, {5, 10}}
	secondList := [][]int{{1, 5}, {8, 12}}
	result := IntervalIntersection(firstList, secondList)
	if len(result) != 3 {
		t.Errorf("expected 3 intersections, got %d", len(result))
	}
}
