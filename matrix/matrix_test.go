package matrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expectedSpiral := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	spiralResult := SpiralOrder(grid)
	if !reflect.DeepEqual(spiralResult, expectedSpiral) {
		t.Errorf("SpiralOrder = %v, want %v", spiralResult, expectedSpiral)
	}
}

func TestSearchSortedMatrix(t *testing.T) {
	grid := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	if !SearchSortedMatrix(grid, 5) {
		t.Error("Should find 5 in matrix")
	}
	if SearchSortedMatrix(grid, 10) {
		t.Error("Should not find 10 in matrix")
	}
}

func TestRotateClockwise90(t *testing.T) {
	grid := [][]int{{1, 2}, {3, 4}}
	RotateClockwise90(grid)
	expectedGrid := [][]int{{3, 1}, {4, 2}}
	if !reflect.DeepEqual(grid, expectedGrid) {
		t.Errorf("RotateClockwise90 = %v, want %v", grid, expectedGrid)
	}
}
