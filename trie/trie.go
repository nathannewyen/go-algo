package trie

// TrieNode represents a single node in the trie.
// Each node stores a map of children (one per character) and a flag
// indicating whether this node marks the end of a complete word.
type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

// Trie implements a prefix tree for efficient string operations.
// Supports insert, search, prefix matching, and autocomplete.
type Trie struct {
	root      *TrieNode
	wordCount int
}

// New creates and returns an empty trie.
func New() *Trie {
	return &Trie{
		root: &TrieNode{
			children:    make(map[rune]*TrieNode),
			isEndOfWord: false,
		},
		wordCount: 0,
	}
}

// Insert adds a word to the trie, creating nodes as needed for each character.
func (t *Trie) Insert(word string) {
	currentNode := t.root

	for _, character := range word {
		// Create a new node for this character if it does not exist
		if _, exists := currentNode.children[character]; !exists {
			currentNode.children[character] = &TrieNode{
				children:    make(map[rune]*TrieNode),
				isEndOfWord: false,
			}
		}
		currentNode = currentNode.children[character]
	}

	// Only increment word count if this is a new word
	if !currentNode.isEndOfWord {
		currentNode.isEndOfWord = true
		t.wordCount++
	}
}

// Search checks if an exact word exists in the trie.
func (t *Trie) Search(word string) bool {
	endNode := t.findNode(word)
	return endNode != nil && endNode.isEndOfWord
}

// StartsWith checks if any word in the trie has the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

// findNode traverses the trie following the characters of the given string.
// Returns the final node if all characters are found, nil otherwise.
func (t *Trie) findNode(text string) *TrieNode {
	currentNode := t.root

	for _, character := range text {
		childNode, exists := currentNode.children[character]
		if !exists {
			return nil
		}
		currentNode = childNode
	}

	return currentNode
}

// AutoComplete returns all words in the trie that start with the given prefix.
func (t *Trie) AutoComplete(prefix string) []string {
	prefixEndNode := t.findNode(prefix)
	if prefixEndNode == nil {
		return nil
	}

	matchingWords := make([]string, 0)
	collectWords(prefixEndNode, prefix, &matchingWords)

	return matchingWords
}

// collectWords recursively gathers all complete words reachable from the given node.
func collectWords(currentNode *TrieNode, currentPrefix string, matchingWords *[]string) {
	if currentNode.isEndOfWord {
		*matchingWords = append(*matchingWords, currentPrefix)
	}

	for character, childNode := range currentNode.children {
		collectWords(childNode, currentPrefix+string(character), matchingWords)
	}
}

// WordCount returns the number of unique words stored in the trie.
func (t *Trie) WordCount() int {
	return t.wordCount
}

// Delete removes a word from the trie. Returns true if the word was found and removed.
func (t *Trie) Delete(word string) bool {
	wordFound, _ := t.deleteRecursive(t.root, []rune(word), 0)
	return wordFound
}

// deleteRecursive walks the trie to find the word, then removes nodes
// on the way back up if they are no longer needed by other words.
// Returns (wordFound, canPruneNode).
func (t *Trie) deleteRecursive(currentNode *TrieNode, wordRunes []rune, depth int) (bool, bool) {
	if depth == len(wordRunes) {
		// Reached the end of the word
		if !currentNode.isEndOfWord {
			return false, false
		}
		currentNode.isEndOfWord = false
		t.wordCount--

		// This node can be pruned if it has no children
		canPrune := len(currentNode.children) == 0
		return true, canPrune
	}

	currentCharacter := wordRunes[depth]
	childNode, exists := currentNode.children[currentCharacter]
	if !exists {
		return false, false
	}

	wordFound, shouldDeleteChild := t.deleteRecursive(childNode, wordRunes, depth+1)

	if shouldDeleteChild {
		delete(currentNode.children, currentCharacter)
		// This node can also be pruned if it has no other children and is not a word end
		canPrune := len(currentNode.children) == 0 && !currentNode.isEndOfWord
		return wordFound, canPrune
	}

	return wordFound, false
}
