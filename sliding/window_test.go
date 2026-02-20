package sliding

import "testing"

func TestMaxSumSubarray(t *testing.T) {
	testCases := []struct {
		description string
		numbers     []int
		windowSize  int
		expectedSum int
		expectedOk  bool
	}{
		{
			description: "basic case",
			numbers:     []int{1, 4, 2, 10, 2, 3, 1, 0, 20},
			windowSize:  4,
			expectedSum: 24,
			expectedOk:  true,
		},
		{
			description: "window equals slice length",
			numbers:     []int{1, 2, 3},
			windowSize:  3,
			expectedSum: 6,
			expectedOk:  true,
		},
		{
			description: "window of size 1",
			numbers:     []int{5, 1, 8, 3},
			windowSize:  1,
			expectedSum: 8,
			expectedOk:  true,
		},
		{
			description: "window larger than slice",
			numbers:     []int{1, 2},
			windowSize:  5,
			expectedSum: 0,
			expectedOk:  false,
		},
		{
			description: "negative numbers",
			numbers:     []int{-1, -2, 5, -1, 3},
			windowSize:  2,
			expectedSum: 4,
			expectedOk:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualSum, actualOk := MaxSumSubarray(testCase.numbers, testCase.windowSize)

			if actualOk != testCase.expectedOk {
				t.Errorf("expected ok=%v, got ok=%v", testCase.expectedOk, actualOk)
			}
			if actualSum != testCase.expectedSum {
				t.Errorf("expected sum %d, got %d", testCase.expectedSum, actualSum)
			}
		})
	}
}

func TestLongestUniqueSubstring(t *testing.T) {
	testCases := []struct {
		description    string
		inputString    string
		expectedLength int
	}{
		{"basic repeating", "abcabcbb", 3},
		{"all same characters", "bbbbb", 1},
		{"mixed unique", "pwwkew", 3},
		{"empty string", "", 0},
		{"all unique", "abcdef", 6},
		{"single character", "a", 1},
		{"repeat at end", "abcda", 4},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualLength := LongestUniqueSubstring(testCase.inputString)

			if actualLength != testCase.expectedLength {
				t.Errorf("expected %d, got %d", testCase.expectedLength, actualLength)
			}
		})
	}
}

func TestMinWindowContaining(t *testing.T) {
	testCases := []struct {
		description      string
		sourceString     string
		targetCharacters string
		expectedWindow   string
		expectedFound    bool
	}{
		{
			description:      "classic example",
			sourceString:     "ADOBECODEBANC",
			targetCharacters: "ABC",
			expectedWindow:   "BANC",
			expectedFound:    true,
		},
		{
			description:      "exact match",
			sourceString:     "ABC",
			targetCharacters: "ABC",
			expectedWindow:   "ABC",
			expectedFound:    true,
		},
		{
			description:      "target not in source",
			sourceString:     "ADOBECODE",
			targetCharacters: "XYZ",
			expectedWindow:   "",
			expectedFound:    false,
		},
		{
			description:      "single character target",
			sourceString:     "hello",
			targetCharacters: "l",
			expectedWindow:   "l",
			expectedFound:    true,
		},
		{
			description:      "empty source",
			sourceString:     "",
			targetCharacters: "ABC",
			expectedWindow:   "",
			expectedFound:    false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualWindow, actualFound := MinWindowContaining(testCase.sourceString, testCase.targetCharacters)

			if actualFound != testCase.expectedFound {
				t.Errorf("expected found=%v, got found=%v", testCase.expectedFound, actualFound)
			}
			if actualWindow != testCase.expectedWindow {
				t.Errorf("expected '%s', got '%s'", testCase.expectedWindow, actualWindow)
			}
		})
	}
}
