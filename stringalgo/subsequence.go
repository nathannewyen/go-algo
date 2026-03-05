package stringalgo

// LongestCommonSubsequence finds the longest subsequence present in both strings
func LongestCommonSubsequence(firstString string, secondString string) string {
	firstLength := len(firstString)
	secondLength := len(secondString)

	// dpTable[i][j] = length of LCS of firstString[:i] and secondString[:j]
	dpTable := make([][]int, firstLength+1)
	for rowIndex := range dpTable {
		dpTable[rowIndex] = make([]int, secondLength+1)
	}

	for firstIndex := 1; firstIndex <= firstLength; firstIndex++ {
		for secondIndex := 1; secondIndex <= secondLength; secondIndex++ {
			if firstString[firstIndex-1] == secondString[secondIndex-1] {
				dpTable[firstIndex][secondIndex] = dpTable[firstIndex-1][secondIndex-1] + 1
			} else {
				topValue := dpTable[firstIndex-1][secondIndex]
				leftValue := dpTable[firstIndex][secondIndex-1]
				if topValue > leftValue {
					dpTable[firstIndex][secondIndex] = topValue
				} else {
					dpTable[firstIndex][secondIndex] = leftValue
				}
			}
		}
	}

	// Backtrack to find the actual subsequence
	subsequenceChars := []byte{}
	firstIndex := firstLength
	secondIndex := secondLength
	for firstIndex > 0 && secondIndex > 0 {
		if firstString[firstIndex-1] == secondString[secondIndex-1] {
			subsequenceChars = append([]byte{firstString[firstIndex-1]}, subsequenceChars...)
			firstIndex--
			secondIndex--
		} else if dpTable[firstIndex-1][secondIndex] > dpTable[firstIndex][secondIndex-1] {
			firstIndex--
		} else {
			secondIndex--
		}
	}
	return string(subsequenceChars)
}
