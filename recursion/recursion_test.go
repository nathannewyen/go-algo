package recursion

import "testing"

func TestGeneratePermutations(t *testing.T) {
	permutations := GeneratePermutations([]int{1, 2, 3})
	if len(permutations) != 6 {
		t.Errorf("expected 6 permutations, got %d", len(permutations))
	}
}

func TestGenerateSubsets(t *testing.T) {
	subsets := GenerateSubsets([]int{1, 2, 3})
	if len(subsets) != 8 {
		t.Errorf("expected 8 subsets, got %d", len(subsets))
	}
}

func TestCombinationSum(t *testing.T) {
	combinations := CombinationSum([]int{2, 3, 6, 7}, 7)
	if len(combinations) != 2 {
		t.Errorf("expected 2 combinations, got %d", len(combinations))
	}
}

func TestLetterCombinations(t *testing.T) {
	combos := LetterCombinations("23")
	if len(combos) != 9 {
		t.Errorf("expected 9 combinations, got %d", len(combos))
	}
}
