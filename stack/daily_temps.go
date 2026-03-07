package stack

// DailyTemperatures returns days to wait for a warmer temperature at each position
func DailyTemperatures(temperatures []int) []int {
	dayCount := len(temperatures)
	daysToWait := make([]int, dayCount)
	// Stack stores indices of temperatures we haven't found a warmer day for
	pendingIndices := []int{}

	for currentDay := 0; currentDay < dayCount; currentDay++ {
		// Pop indices where current temperature is warmer than the pending day
		for len(pendingIndices) > 0 {
			topPendingIndex := pendingIndices[len(pendingIndices)-1]
			if temperatures[currentDay] <= temperatures[topPendingIndex] {
				break
			}
			pendingIndices = pendingIndices[:len(pendingIndices)-1]
			daysToWait[topPendingIndex] = currentDay - topPendingIndex
		}
		pendingIndices = append(pendingIndices, currentDay)
	}
	return daysToWait
}
