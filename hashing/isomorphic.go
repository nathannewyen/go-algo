package hashing

// AreIsomorphic checks if two strings have the same character mapping pattern
func AreIsomorphic(firstString string, secondString string) bool {
	if len(firstString) != len(secondString) {
		return false
	}
	// Track bidirectional character mappings
	firstToSecondMapping := map[byte]byte{}
	secondToFirstMapping := map[byte]byte{}

	for i := 0; i < len(firstString); i++ {
		charFromFirst := firstString[i]
		charFromSecond := secondString[i]
		// Check if existing mapping from first to second is consistent
		if mappedChar, exists := firstToSecondMapping[charFromFirst]; exists {
			if mappedChar != charFromSecond {
				return false
			}
		} else {
			firstToSecondMapping[charFromFirst] = charFromSecond
		}
		// Check reverse mapping consistency
		if mappedChar, exists := secondToFirstMapping[charFromSecond]; exists {
			if mappedChar != charFromFirst {
				return false
			}
		} else {
			secondToFirstMapping[charFromSecond] = charFromFirst
		}
	}
	return true
}
