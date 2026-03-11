package monotonic

// TrappingRainWater calculates total water trapped between elevation bars
func TrappingRainWater(elevationMap []int) int {
	barCount := len(elevationMap)
	if barCount < 3 {
		return 0
	}
	// Precompute the maximum height to the left of each position
	maxHeightFromLeft := make([]int, barCount)
	maxHeightFromLeft[0] = elevationMap[0]
	for i := 1; i < barCount; i++ {
		if elevationMap[i] > maxHeightFromLeft[i-1] {
			maxHeightFromLeft[i] = elevationMap[i]
		} else {
			maxHeightFromLeft[i] = maxHeightFromLeft[i-1]
		}
	}
	// Precompute the maximum height to the right of each position
	maxHeightFromRight := make([]int, barCount)
	maxHeightFromRight[barCount-1] = elevationMap[barCount-1]
	for i := barCount - 2; i >= 0; i-- {
		if elevationMap[i] > maxHeightFromRight[i+1] {
			maxHeightFromRight[i] = elevationMap[i]
		} else {
			maxHeightFromRight[i] = maxHeightFromRight[i+1]
		}
	}
	// Water at each position is bounded by the shorter of the two max heights
	totalTrappedWater := 0
	for i := 1; i < barCount-1; i++ {
		boundingHeight := maxHeightFromLeft[i]
		if maxHeightFromRight[i] < boundingHeight {
			boundingHeight = maxHeightFromRight[i]
		}
		waterAtPosition := boundingHeight - elevationMap[i]
		if waterAtPosition > 0 {
			totalTrappedWater += waterAtPosition
		}
	}
	return totalTrappedWater
}
