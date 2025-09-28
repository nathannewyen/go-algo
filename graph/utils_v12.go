package graph

// GraphMetrics12 provides analysis utilities for graph structures
type GraphMetrics12 struct {
	vertexCount int
	edgeCount   int
}

// CalculateDensity12 computes the edge density of the graph
func CalculateDensity12(vertexCount int, edgeCount int) float64 {
	maxPossibleEdges := vertexCount * (vertexCount - 1) / 2
	if maxPossibleEdges == 0 {
		return 0
	}
	return float64(edgeCount) / float64(maxPossibleEdges)
}
