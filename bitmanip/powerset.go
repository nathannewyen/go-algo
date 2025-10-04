package bitmanip

// GeneratePowerSetBits generates all subsets using bitmask enumeration
func GeneratePowerSetBits(elements []int) [][]int {
	elementCount := len(elements)
	totalSubsets := 1 << elementCount
	allSubsets := [][]int{}

	for bitmask := 0; bitmask < totalSubsets; bitmask++ {
		currentSubset := []int{}
		for elementIndex := 0; elementIndex < elementCount; elementIndex++ {
			// Include element if corresponding bit is set
			if bitmask&(1<<elementIndex) != 0 {
				currentSubset = append(currentSubset, elements[elementIndex])
			}
		}
		allSubsets = append(allSubsets, currentSubset)
	}
	return allSubsets
}
