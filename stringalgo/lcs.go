package stringalgo

// LongestCommonSubstring finds the longest substring shared between two strings
func LongestCommonSubstring(firstString string, secondString string) string {
	firstLength := len(firstString)
	secondLength := len(secondString)

	// DP table where dpTable[i][j] = length of common substring ending at i,j
	dpTable := make([][]int, firstLength+1)
	for rowIndex := range dpTable {
		dpTable[rowIndex] = make([]int, secondLength+1)
	}

	maxSubstringLength := 0
	endPositionInFirst := 0

	for firstIndex := 1; firstIndex <= firstLength; firstIndex++ {
		for secondIndex := 1; secondIndex <= secondLength; secondIndex++ {
			if firstString[firstIndex-1] == secondString[secondIndex-1] {
				dpTable[firstIndex][secondIndex] = dpTable[firstIndex-1][secondIndex-1] + 1
				if dpTable[firstIndex][secondIndex] > maxSubstringLength {
					maxSubstringLength = dpTable[firstIndex][secondIndex]
					endPositionInFirst = firstIndex
				}
			}
		}
	}

	if maxSubstringLength == 0 {
		return ""
	}
	return firstString[endPositionInFirst-maxSubstringLength : endPositionInFirst]
}
