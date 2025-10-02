package greedy

// FindStartingGasStation finds the starting station to complete a circular tour
// Returns -1 if no valid starting station exists
func FindStartingGasStation(gasAmounts []int, travelCosts []int) int {
	totalSurplus := 0
	currentTankLevel := 0
	startingStation := 0

	for stationIndex := 0; stationIndex < len(gasAmounts); stationIndex++ {
		netGainAtStation := gasAmounts[stationIndex] - travelCosts[stationIndex]
		totalSurplus += netGainAtStation
		currentTankLevel += netGainAtStation

		// If tank goes negative, reset start to next station
		if currentTankLevel < 0 {
			startingStation = stationIndex + 1
			currentTankLevel = 0
		}
	}

	// Circuit is possible only if total gas >= total cost
	if totalSurplus >= 0 {
		return startingStation
	}
	return -1
}
