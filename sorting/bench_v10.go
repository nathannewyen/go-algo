package sorting

// SortingBenchmark10 tracks performance of sorting algorithms
type SortingBenchmark10 struct {
	algorithmName  string
	comparisonCount int
	swapCount       int
}

// NewBenchmark10 creates a new benchmark tracker
func NewBenchmark10(algorithmName string) *SortingBenchmark10 {
	return &SortingBenchmark10{algorithmName: algorithmName}
}

// RecordComparison10 increments the comparison counter
func (benchmark *SortingBenchmark10) RecordComparison10() {
	benchmark.comparisonCount++
}

// RecordSwap10 increments the swap counter
func (benchmark *SortingBenchmark10) RecordSwap10() {
	benchmark.swapCount++
}
