package intervals

// CanAttendAllMeetings checks if a person can attend all meetings without overlap
func CanAttendAllMeetings(meetings [][]int) bool {
	sortIntervals(meetings)
	for i := 1; i < len(meetings); i++ {
		previousEndTime := meetings[i-1][1]
		currentStartTime := meetings[i][0]
		// If current meeting starts before previous one ends, conflict exists
		if currentStartTime < previousEndTime {
			return false
		}
	}
	return true
}

// MinMeetingRooms returns the minimum number of conference rooms required
func MinMeetingRooms(meetings [][]int) int {
	if len(meetings) == 0 {
		return 0
	}
	// Collect all start and end times as events
	startTimes := make([]int, len(meetings))
	endTimes := make([]int, len(meetings))
	for i, meeting := range meetings {
		startTimes[i] = meeting[0]
		endTimes[i] = meeting[1]
	}
	sortSlice(startTimes)
	sortSlice(endTimes)
	roomsNeeded := 0
	maxRoomsNeeded := 0
	endPointer := 0
	// Sweep through start times and track active meetings
	for startPointer := 0; startPointer < len(startTimes); startPointer++ {
		if startTimes[startPointer] < endTimes[endPointer] {
			roomsNeeded++
		} else {
			endPointer++
		}
		if roomsNeeded > maxRoomsNeeded {
			maxRoomsNeeded = roomsNeeded
		}
	}
	return maxRoomsNeeded
}

// sortSlice sorts an integer slice using insertion sort
func sortSlice(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}
