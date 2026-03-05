package stringalgo

import "testing"

func TestKMPSearch(t *testing.T) {
	matchPositions := KMPSearch("AABAACAADAABAABA", "AABA")
	if len(matchPositions) != 3 {
		t.Errorf("Expected 3 matches, got %d", len(matchPositions))
	}
}

func TestRabinKarpSearch(t *testing.T) {
	matchPositions := RabinKarpSearch("hello world hello", "hello")
	if len(matchPositions) != 2 {
		t.Errorf("Expected 2 matches, got %d", len(matchPositions))
	}
}

func TestLongestCommonSubstring(t *testing.T) {
	result := LongestCommonSubstring("ABABC", "BABCBA")
	if result != "BABC" {
		t.Errorf("Expected BABC, got %s", result)
	}
}

func TestLongestPalindromicSubstring(t *testing.T) {
	result := LongestPalindromicSubstring("babad")
	if result != "bab" && result != "aba" {
		t.Errorf("Expected bab or aba, got %s", result)
	}
}

func TestAreAnagrams(t *testing.T) {
	if !AreAnagrams("listen", "silent") {
		t.Error("listen and silent should be anagrams")
	}
}
