package bitmanip

import "testing"

func TestCountSetBits(t *testing.T) {
	if CountSetBits(11) != 3 {
		t.Errorf("CountSetBits(11) should be 3, got %d", CountSetBits(11))
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	if !IsPowerOfTwo(16) {
		t.Error("16 should be power of two")
	}
	if IsPowerOfTwo(15) {
		t.Error("15 should not be power of two")
	}
}

func TestSingleNumber(t *testing.T) {
	result := SingleNumber([]int{4, 1, 2, 1, 2})
	if result != 4 {
		t.Errorf("SingleNumber should return 4, got %d", result)
	}
}

func TestHammingDistance(t *testing.T) {
	if HammingDistance(1, 4) != 2 {
		t.Errorf("HammingDistance(1,4) should be 2")
	}
}

func TestFindMissingNumber(t *testing.T) {
	result := FindMissingNumber([]int{3, 0, 1})
	if result != 2 {
		t.Errorf("Missing number should be 2, got %d", result)
	}
}
