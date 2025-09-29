package sorting

// SortingBenchmark11 tracks performance of sorting algorithms
type SortingBenchmark11 struct {
	algorithmName  string
	comparisonCount int
	swapCount       int
}

// NewBenchmark11 creates a new benchmark tracker
func NewBenchmark11(algorithmName string) *SortingBenchmark11 {
	return &SortingBenchmark11{algorithmName: algorithmName}
}

// RecordComparison11 increments the comparison counter
func (benchmark *SortingBenchmark11) RecordComparison11() {
	benchmark.comparisonCount++
}

// RecordSwap11 increments the swap counter
func (benchmark *SortingBenchmark11) RecordSwap11() {
	benchmark.swapCount++
}
