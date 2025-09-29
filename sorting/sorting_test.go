package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	unsortedNumbers := []int{5, 3, 8, 1, 2}
	expectedResult := []int{1, 2, 3, 5, 8}
	sortedResult := BubbleSort(unsortedNumbers)
	if !reflect.DeepEqual(sortedResult, expectedResult) {
		t.Errorf("BubbleSort = %v, want %v", sortedResult, expectedResult)
	}
}

func TestInsertionSort(t *testing.T) {
	unsortedNumbers := []int{5, 3, 8, 1, 2}
	expectedResult := []int{1, 2, 3, 5, 8}
	sortedResult := InsertionSort(unsortedNumbers)
	if !reflect.DeepEqual(sortedResult, expectedResult) {
		t.Errorf("InsertionSort = %v, want %v", sortedResult, expectedResult)
	}
}

func TestMergeSort(t *testing.T) {
	unsortedNumbers := []int{5, 3, 8, 1, 2}
	expectedResult := []int{1, 2, 3, 5, 8}
	sortedResult := MergeSort(unsortedNumbers)
	if !reflect.DeepEqual(sortedResult, expectedResult) {
		t.Errorf("MergeSort = %v, want %v", sortedResult, expectedResult)
	}
}

func TestQuickSort(t *testing.T) {
	unsortedNumbers := []int{5, 3, 8, 1, 2}
	expectedResult := []int{1, 2, 3, 5, 8}
	sortedResult := QuickSort(unsortedNumbers)
	if !reflect.DeepEqual(sortedResult, expectedResult) {
		t.Errorf("QuickSort = %v, want %v", sortedResult, expectedResult)
	}
}

func TestHeapSort(t *testing.T) {
	unsortedNumbers := []int{5, 3, 8, 1, 2}
	expectedResult := []int{1, 2, 3, 5, 8}
	sortedResult := HeapSort(unsortedNumbers)
	if !reflect.DeepEqual(sortedResult, expectedResult) {
		t.Errorf("HeapSort = %v, want %v", sortedResult, expectedResult)
	}
}
