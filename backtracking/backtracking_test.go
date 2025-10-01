package backtracking

import "testing"

func TestGeneratePermutations(t *testing.T) {
	permutationResults := GeneratePermutations([]int{1, 2, 3})
	// 3! = 6 permutations expected
	if len(permutationResults) != 6 {
		t.Errorf("Expected 6 permutations, got %d", len(permutationResults))
	}
}

func TestGenerateCombinations(t *testing.T) {
	combinationResults := GenerateCombinations(4, 2)
	// C(4,2) = 6 combinations expected
	if len(combinationResults) != 6 {
		t.Errorf("Expected 6 combinations, got %d", len(combinationResults))
	}
}

func TestGenerateSubsets(t *testing.T) {
	subsetResults := GenerateSubsets([]int{1, 2, 3})
	// 2^3 = 8 subsets expected
	if len(subsetResults) != 8 {
		t.Errorf("Expected 8 subsets, got %d", len(subsetResults))
	}
}

func TestSolveNQueens(t *testing.T) {
	solutionResults := SolveNQueens(4)
	// 4-Queens has 2 solutions
	if len(solutionResults) != 2 {
		t.Errorf("Expected 2 solutions for 4-Queens, got %d", len(solutionResults))
	}
}
