package hashing

// GroupAnagrams groups strings that are anagrams of each other
func GroupAnagrams(words []string) [][]string {
	// Use sorted character frequency as the grouping key
	anagramGroups := map[string][]string{}
	for _, word := range words {
		sortedKey := buildSortedKey(word)
		anagramGroups[sortedKey] = append(anagramGroups[sortedKey], word)
	}
	groupedResults := [][]string{}
	for _, group := range anagramGroups {
		groupedResults = append(groupedResults, group)
	}
	return groupedResults
}

// buildSortedKey creates a canonical key from character frequencies
func buildSortedKey(word string) string {
	charFrequency := [26]int{}
	for i := 0; i < len(word); i++ {
		charFrequency[word[i]-'a']++
	}
	keyBytes := []byte{}
	for i := 0; i < 26; i++ {
		for j := 0; j < charFrequency[i]; j++ {
			keyBytes = append(keyBytes, byte(i)+'a')
		}
	}
	return string(keyBytes)
}
