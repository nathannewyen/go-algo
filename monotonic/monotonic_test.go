package monotonic

import "testing"

func TestNextSmallerElement(t *testing.T) {
	numbers := []int{4, 8, 5, 2, 25}
	result := NextSmallerElement(numbers)
	if result[0] != 2 {
		t.Errorf("expected next smaller of 4 to be 2, got %d", result[0])
	}
}

func TestLargestRectangle(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	area := LargestRectangleInHistogram(heights)
	if area != 10 {
		t.Errorf("expected area 10, got %d", area)
	}
}

func TestTrappingRainWater(t *testing.T) {
	elevation := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trapped := TrappingRainWater(elevation)
	if trapped != 6 {
		t.Errorf("expected 6 units of water, got %d", trapped)
	}
}

func TestStockSpan(t *testing.T) {
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	spans := StockSpan(prices)
	expected := []int{1, 1, 1, 2, 1, 4, 6}
	for i := range expected {
		if spans[i] != expected[i] {
			t.Errorf("day %d: expected span %d, got %d", i, expected[i], spans[i])
		}
	}
}
