package search

import "testing"

func TestBinarySearch(t *testing.T) {
	sortedNumbers := []int{1, 3, 5, 7, 9, 11, 13, 15}

	testCases := []struct {
		description   string
		targetValue   int
		expectedIndex int
		expectedFound bool
	}{
		{"find first element", 1, 0, true},
		{"find middle element", 7, 3, true},
		{"find last element", 15, 7, true},
		{"find element at index 5", 11, 5, true},
		{"missing value between elements", 6, -1, false},
		{"missing value below range", 0, -1, false},
		{"missing value above range", 20, -1, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualIndex, actualFound := BinarySearch(sortedNumbers, testCase.targetValue)

			if actualIndex != testCase.expectedIndex {
				t.Errorf("expected index %d, got %d", testCase.expectedIndex, actualIndex)
			}
			if actualFound != testCase.expectedFound {
				t.Errorf("expected found %v, got %v", testCase.expectedFound, actualFound)
			}
		})
	}
}

func TestBinarySearchEmptySlice(t *testing.T) {
	emptySlice := []int{}

	_, found := BinarySearch(emptySlice, 5)
	if found {
		t.Error("expected not found in empty slice")
	}
}

func TestBinarySearchStrings(t *testing.T) {
	sortedWords := []string{"apple", "banana", "cherry", "date", "elderberry"}

	index, found := BinarySearch(sortedWords, "cherry")
	if !found {
		t.Error("expected to find 'cherry'")
	}
	if index != 2 {
		t.Errorf("expected index 2, got %d", index)
	}
}

func TestBinarySearchFirstOccurrence(t *testing.T) {
	// Slice with duplicate values
	sortedWithDuplicates := []int{1, 3, 3, 3, 5, 7, 7, 9}

	testCases := []struct {
		description   string
		targetValue   int
		expectedIndex int
		expectedFound bool
	}{
		{"first occurrence of 3", 3, 1, true},
		{"first occurrence of 7", 7, 5, true},
		{"unique value 1", 1, 0, true},
		{"unique value 9", 9, 7, true},
		{"missing value", 4, -1, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualIndex, actualFound := BinarySearchFirstOccurrence(sortedWithDuplicates, testCase.targetValue)

			if actualIndex != testCase.expectedIndex {
				t.Errorf("expected index %d, got %d", testCase.expectedIndex, actualIndex)
			}
			if actualFound != testCase.expectedFound {
				t.Errorf("expected found %v, got %v", testCase.expectedFound, actualFound)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	sortedNumbers := []int{1, 2, 4, 7, 11, 15}

	testCases := []struct {
		description        string
		targetSum          int
		expectedLeftIndex  int
		expectedRightIndex int
		expectedFound      bool
	}{
		{"sum to 9", 9, 1, 3, true},
		{"sum to 15", 15, 2, 4, true},
		{"sum to 3", 3, 0, 1, true},
		{"impossible sum", 100, -1, -1, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualLeft, actualRight, actualFound := TwoSum(sortedNumbers, testCase.targetSum)

			if actualLeft != testCase.expectedLeftIndex {
				t.Errorf("expected left index %d, got %d", testCase.expectedLeftIndex, actualLeft)
			}
			if actualRight != testCase.expectedRightIndex {
				t.Errorf("expected right index %d, got %d", testCase.expectedRightIndex, actualRight)
			}
			if actualFound != testCase.expectedFound {
				t.Errorf("expected found %v, got %v", testCase.expectedFound, actualFound)
			}
		})
	}
}
