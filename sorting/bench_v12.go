package sorting

// SortingBenchmark12 tracks performance of sorting algorithms
type SortingBenchmark12 struct {
	algorithmName  string
	comparisonCount int
	swapCount       int
}

// NewBenchmark12 creates a new benchmark tracker
func NewBenchmark12(algorithmName string) *SortingBenchmark12 {
	return &SortingBenchmark12{algorithmName: algorithmName}
}

// RecordComparison12 increments the comparison counter
func (benchmark *SortingBenchmark12) RecordComparison12() {
	benchmark.comparisonCount++
}

// RecordSwap12 increments the swap counter
func (benchmark *SortingBenchmark12) RecordSwap12() {
	benchmark.swapCount++
}
