package hashing

import "testing"

func TestTwoSum(t *testing.T) {
	result := TwoSum([]int{2, 7, 11, 15}, 9)
	if len(result) != 2 {
		t.Fatalf("expected 2 indices, got %d", len(result))
	}
	if result[0] != 0 {
		t.Errorf("expected first index 0, got %d", result[0])
	}
}

func TestGroupAnagrams(t *testing.T) {
	groups := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	if len(groups) != 3 {
		t.Errorf("expected 3 groups, got %d", len(groups))
	}
}

func TestLongestConsecutive(t *testing.T) {
	length := LongestConsecutiveSequence([]int{100, 4, 200, 1, 3, 2})
	if length != 4 {
		t.Errorf("expected longest streak 4, got %d", length)
	}
}

func TestTopKFrequent(t *testing.T) {
	result := TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	if len(result) != 2 {
		t.Errorf("expected 2 elements, got %d", len(result))
	}
}

func TestContainsDuplicate(t *testing.T) {
	if !ContainsDuplicate([]int{1, 2, 3, 1}) {
		t.Error("expected duplicate found")
	}
	if ContainsDuplicate([]int{1, 2, 3, 4}) {
		t.Error("expected no duplicate")
	}
}
