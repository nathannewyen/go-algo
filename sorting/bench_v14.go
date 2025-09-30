package sorting

// SortingBenchmark14 tracks performance of sorting algorithms
type SortingBenchmark14 struct {
	algorithmName  string
	comparisonCount int
	swapCount       int
}

// NewBenchmark14 creates a new benchmark tracker
func NewBenchmark14(algorithmName string) *SortingBenchmark14 {
	return &SortingBenchmark14{algorithmName: algorithmName}
}

// RecordComparison14 increments the comparison counter
func (benchmark *SortingBenchmark14) RecordComparison14() {
	benchmark.comparisonCount++
}

// RecordSwap14 increments the swap counter
func (benchmark *SortingBenchmark14) RecordSwap14() {
	benchmark.swapCount++
}
