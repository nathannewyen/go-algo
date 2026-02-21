package dp

import "testing"

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		nthPosition    int
		expectedResult int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{20, 6765},
	}

	for _, testCase := range testCases {
		actualResult := Fibonacci(testCase.nthPosition)
		if actualResult != testCase.expectedResult {
			t.Errorf("Fibonacci(%d) = %d, expected %d", testCase.nthPosition, actualResult, testCase.expectedResult)
		}
	}
}

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		totalStairs    int
		expectedResult int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, testCase := range testCases {
		actualResult := ClimbStairs(testCase.totalStairs)
		if actualResult != testCase.expectedResult {
			t.Errorf("ClimbStairs(%d) = %d, expected %d", testCase.totalStairs, actualResult, testCase.expectedResult)
		}
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		description    string
		firstString    string
		secondString   string
		expectedLength int
	}{
		{"basic case", "abcde", "ace", 3},
		{"identical strings", "abc", "abc", 3},
		{"no common subsequence", "abc", "xyz", 0},
		{"empty first string", "", "abc", 0},
		{"empty second string", "abc", "", 0},
		{"single char match", "a", "a", 1},
		{"longer example", "AGGTAB", "GXTXAYB", 4},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualLength := LongestCommonSubsequence(testCase.firstString, testCase.secondString)
			if actualLength != testCase.expectedLength {
				t.Errorf("expected %d, got %d", testCase.expectedLength, actualLength)
			}
		})
	}
}

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		description      string
		coinDenominations []int
		targetAmount     int
		expectedMinCoins int
	}{
		{"standard coins", []int{1, 5, 10, 25}, 30, 2},
		{"exact coin match", []int{1, 5, 10}, 10, 1},
		{"impossible amount", []int{2}, 3, -1},
		{"zero amount", []int{1}, 0, 0},
		{"single denomination", []int{3}, 9, 3},
		{"multiple options", []int{1, 3, 4}, 6, 2},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualMinCoins := CoinChange(testCase.coinDenominations, testCase.targetAmount)
			if actualMinCoins != testCase.expectedMinCoins {
				t.Errorf("expected %d, got %d", testCase.expectedMinCoins, actualMinCoins)
			}
		})
	}
}

func TestMaxSubarraySum(t *testing.T) {
	testCases := []struct {
		description    string
		numbers        []int
		expectedMaxSum int
	}{
		{"mixed positive and negative", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"all positive", []int{1, 2, 3, 4}, 10},
		{"all negative", []int{-1, -2, -3}, -1},
		{"single element", []int{5}, 5},
		{"empty slice", []int{}, 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualMaxSum := MaxSubarraySum(testCase.numbers)
			if actualMaxSum != testCase.expectedMaxSum {
				t.Errorf("expected %d, got %d", testCase.expectedMaxSum, actualMaxSum)
			}
		})
	}
}
