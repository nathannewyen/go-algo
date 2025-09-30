package sorting

// SortingBenchmark13 tracks performance of sorting algorithms
type SortingBenchmark13 struct {
	algorithmName  string
	comparisonCount int
	swapCount       int
}

// NewBenchmark13 creates a new benchmark tracker
func NewBenchmark13(algorithmName string) *SortingBenchmark13 {
	return &SortingBenchmark13{algorithmName: algorithmName}
}

// RecordComparison13 increments the comparison counter
func (benchmark *SortingBenchmark13) RecordComparison13() {
	benchmark.comparisonCount++
}

// RecordSwap13 increments the swap counter
func (benchmark *SortingBenchmark13) RecordSwap13() {
	benchmark.swapCount++
}
