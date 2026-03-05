package stringalgo

import "sort"

// GroupAnagramsTogether groups strings that are anagrams of each other
func GroupAnagramsTogether(wordList []string) [][]string {
	// Map sorted characters to group of original words
	anagramGroups := make(map[string][]string)

	for _, currentWord := range wordList {
		sortedKey := sortCharacters(currentWord)
		anagramGroups[sortedKey] = append(anagramGroups[sortedKey], currentWord)
	}

	groupedResult := [][]string{}
	for _, wordGroup := range anagramGroups {
		groupedResult = append(groupedResult, wordGroup)
	}
	return groupedResult
}

// sortCharacters returns a string with characters sorted alphabetically
func sortCharacters(word string) string {
	characterSlice := []byte(word)
	sort.Slice(characterSlice, func(i int, j int) bool {
		return characterSlice[i] < characterSlice[j]
	})
	return string(characterSlice)
}

// AreAnagrams checks if two strings are anagrams of each other
func AreAnagrams(firstWord string, secondWord string) bool {
	if len(firstWord) != len(secondWord) {
		return false
	}
	// Count character frequencies and compare
	charFrequency := make(map[rune]int)
	for _, character := range firstWord {
		charFrequency[character]++
	}
	for _, character := range secondWord {
		charFrequency[character]--
		if charFrequency[character] < 0 {
			return false
		}
	}
	return true
}
