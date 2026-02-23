package trie

import (
	"sort"
	"testing"
)

func TestNewTrie(t *testing.T) {
	prefixTree := New()

	if prefixTree.WordCount() != 0 {
		t.Errorf("expected word count 0, got %d", prefixTree.WordCount())
	}
}

func TestInsertAndSearch(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("apple")
	prefixTree.Insert("app")
	prefixTree.Insert("banana")

	if !prefixTree.Search("apple") {
		t.Error("expected to find 'apple'")
	}
	if !prefixTree.Search("app") {
		t.Error("expected to find 'app'")
	}
	if !prefixTree.Search("banana") {
		t.Error("expected to find 'banana'")
	}

	// Partial words should not be found as exact matches
	if prefixTree.Search("ap") {
		t.Error("did not expect to find 'ap' as an exact word")
	}
	if prefixTree.Search("ban") {
		t.Error("did not expect to find 'ban' as an exact word")
	}
}

func TestInsertDuplicate(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("hello")
	prefixTree.Insert("hello")

	// Duplicate insertion should not increase word count
	if prefixTree.WordCount() != 1 {
		t.Errorf("expected word count 1 after duplicate insert, got %d", prefixTree.WordCount())
	}
}

func TestStartsWith(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("apple")
	prefixTree.Insert("application")

	if !prefixTree.StartsWith("app") {
		t.Error("expected 'app' to be a valid prefix")
	}
	if !prefixTree.StartsWith("apple") {
		t.Error("expected 'apple' to be a valid prefix")
	}
	if prefixTree.StartsWith("banana") {
		t.Error("did not expect 'banana' to be a valid prefix")
	}
}

func TestAutoComplete(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("app")
	prefixTree.Insert("apple")
	prefixTree.Insert("application")
	prefixTree.Insert("banana")

	// AutoComplete should return all words starting with the prefix
	completions := prefixTree.AutoComplete("app")
	sort.Strings(completions)

	expectedCompletions := []string{"app", "apple", "application"}
	sort.Strings(expectedCompletions)

	if len(completions) != len(expectedCompletions) {
		t.Errorf("expected %d completions, got %d", len(expectedCompletions), len(completions))
	}

	for completionIndex, expectedWord := range expectedCompletions {
		if completions[completionIndex] != expectedWord {
			t.Errorf("expected '%s' at index %d, got '%s'", expectedWord, completionIndex, completions[completionIndex])
		}
	}
}

func TestAutoCompleteNoMatch(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("hello")

	completions := prefixTree.AutoComplete("xyz")
	if completions != nil {
		t.Errorf("expected nil for non-matching prefix, got %v", completions)
	}
}

func TestDelete(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("apple")
	prefixTree.Insert("app")

	// Delete "apple" — "app" should still exist
	removed := prefixTree.Delete("apple")
	if !removed {
		t.Error("expected delete to return true for existing word")
	}
	if prefixTree.Search("apple") {
		t.Error("expected 'apple' to be removed")
	}
	if !prefixTree.Search("app") {
		t.Error("expected 'app' to still exist after deleting 'apple'")
	}
	if prefixTree.WordCount() != 1 {
		t.Errorf("expected word count 1, got %d", prefixTree.WordCount())
	}
}

func TestDeleteNonExistent(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("hello")

	removed := prefixTree.Delete("world")
	if removed {
		t.Error("expected delete to return false for non-existent word")
	}
}

func TestDeletePrefixOfAnotherWord(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("app")
	prefixTree.Insert("apple")

	// Delete "app" — "apple" should still be reachable
	prefixTree.Delete("app")

	if prefixTree.Search("app") {
		t.Error("expected 'app' to be removed")
	}
	if !prefixTree.Search("apple") {
		t.Error("expected 'apple' to still exist after deleting 'app'")
	}
}

func TestEmptyString(t *testing.T) {
	prefixTree := New()

	prefixTree.Insert("")

	if !prefixTree.Search("") {
		t.Error("expected to find empty string")
	}
	if prefixTree.WordCount() != 1 {
		t.Errorf("expected word count 1, got %d", prefixTree.WordCount())
	}
}
