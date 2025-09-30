package sorting

// SortingBenchmark15 tracks performance of sorting algorithms
type SortingBenchmark15 struct {
	algorithmName  string
	comparisonCount int
	swapCount       int
}

// NewBenchmark15 creates a new benchmark tracker
func NewBenchmark15(algorithmName string) *SortingBenchmark15 {
	return &SortingBenchmark15{algorithmName: algorithmName}
}

// RecordComparison15 increments the comparison counter
func (benchmark *SortingBenchmark15) RecordComparison15() {
	benchmark.comparisonCount++
}

// RecordSwap15 increments the swap counter
func (benchmark *SortingBenchmark15) RecordSwap15() {
	benchmark.swapCount++
}
