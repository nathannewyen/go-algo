package greedy

import "sort"

// MinPlatformsNeeded finds minimum platforms needed so no train waits
func MinPlatformsNeeded(arrivalTimes []int, departureTimes []int) int {
	sort.Ints(arrivalTimes)
	sort.Ints(departureTimes)

	platformsNeeded := 0
	maximumPlatforms := 0
	arrivalIndex := 0
	departureIndex := 0

	for arrivalIndex < len(arrivalTimes) {
		// If next event is an arrival, need one more platform
		if arrivalTimes[arrivalIndex] <= departureTimes[departureIndex] {
			platformsNeeded++
			if platformsNeeded > maximumPlatforms {
				maximumPlatforms = platformsNeeded
			}
			arrivalIndex++
		} else {
			// A train departed, free up a platform
			platformsNeeded--
			departureIndex++
		}
	}
	return maximumPlatforms
}
