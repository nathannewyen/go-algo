package stringalgo

// CalculateEditDistance computes minimum operations to transform source into target
// Operations: insert, delete, replace (each costs 1)
func CalculateEditDistance(sourceWord string, targetWord string) int {
	sourceLength := len(sourceWord)
	targetLength := len(targetWord)

	// dpTable[i][j] = edit distance between source[:i] and target[:j]
	dpTable := make([][]int, sourceLength+1)
	for rowIndex := range dpTable {
		dpTable[rowIndex] = make([]int, targetLength+1)
	}

	// Base cases: transforming empty string to prefix
	for sourceIndex := 0; sourceIndex <= sourceLength; sourceIndex++ {
		dpTable[sourceIndex][0] = sourceIndex
	}
	for targetIndex := 0; targetIndex <= targetLength; targetIndex++ {
		dpTable[0][targetIndex] = targetIndex
	}

	for sourceIndex := 1; sourceIndex <= sourceLength; sourceIndex++ {
		for targetIndex := 1; targetIndex <= targetLength; targetIndex++ {
			if sourceWord[sourceIndex-1] == targetWord[targetIndex-1] {
				// Characters match, no operation needed
				dpTable[sourceIndex][targetIndex] = dpTable[sourceIndex-1][targetIndex-1]
			} else {
				// Take minimum of insert, delete, replace
				insertCost := dpTable[sourceIndex][targetIndex-1] + 1
				deleteCost := dpTable[sourceIndex-1][targetIndex] + 1
				replaceCost := dpTable[sourceIndex-1][targetIndex-1] + 1

				minimumCost := insertCost
				if deleteCost < minimumCost {
					minimumCost = deleteCost
				}
				if replaceCost < minimumCost {
					minimumCost = replaceCost
				}
				dpTable[sourceIndex][targetIndex] = minimumCost
			}
		}
	}
	return dpTable[sourceLength][targetLength]
}
