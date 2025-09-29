package sorting

// SortingBenchmark9 tracks performance of sorting algorithms
type SortingBenchmark9 struct {
	algorithmName  string
	comparisonCount int
	swapCount       int
}

// NewBenchmark9 creates a new benchmark tracker
func NewBenchmark9(algorithmName string) *SortingBenchmark9 {
	return &SortingBenchmark9{algorithmName: algorithmName}
}

// RecordComparison9 increments the comparison counter
func (benchmark *SortingBenchmark9) RecordComparison9() {
	benchmark.comparisonCount++
}

// RecordSwap9 increments the swap counter
func (benchmark *SortingBenchmark9) RecordSwap9() {
	benchmark.swapCount++
}
