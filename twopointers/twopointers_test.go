package twopointers

import "testing"

func TestFindPairWithTargetSum(t *testing.T) {
	sortedNumbers := []int{1, 2, 3, 4, 6}
	pairResult := FindPairWithTargetSum(sortedNumbers, 6)
	if pairResult[0] == -1 {
		t.Error("Should find pair with sum 6")
	}
}

func TestRemoveDuplicates(t *testing.T) {
	sortedNumbers := []int{1, 1, 2, 3, 3, 4}
	uniqueCount := RemoveDuplicates(sortedNumbers)
	if uniqueCount != 4 {
		t.Errorf("Expected 4 unique elements, got %d", uniqueCount)
	}
}

func TestIsValidPalindrome(t *testing.T) {
	if !IsValidPalindrome("A man, a plan, a canal: Panama") {
		t.Error("Should be a valid palindrome")
	}
	if IsValidPalindrome("hello") {
		t.Error("hello should not be a palindrome")
	}
}
