package prefixsum

import "testing"

func TestBuildPrefixSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	prefixSumArray := BuildPrefixSum(numbers)
	rangeTotal := RangeSum(prefixSumArray, 1, 3)
	if rangeTotal != 9 {
		t.Errorf("expected range sum 9, got %d", rangeTotal)
	}
}

func TestSubarraySumEqualsK(t *testing.T) {
	numbers := []int{1, 1, 1}
	count := SubarraySumEqualsK(numbers, 2)
	if count != 2 {
		t.Errorf("expected 2 subarrays, got %d", count)
	}
}

func TestProductExceptSelf(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	result := ProductExceptSelf(numbers)
	expected := []int{24, 12, 8, 6}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("index %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}

func TestMaxContiguousSum(t *testing.T) {
	numbers := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := MaxContiguousSum(numbers)
	if maxSum != 6 {
		t.Errorf("expected 6, got %d", maxSum)
	}
}
