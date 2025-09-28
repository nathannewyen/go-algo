package graph

// GraphMetrics10 provides analysis utilities for graph structures
type GraphMetrics10 struct {
	vertexCount int
	edgeCount   int
}

// CalculateDensity10 computes the edge density of the graph
func CalculateDensity10(vertexCount int, edgeCount int) float64 {
	maxPossibleEdges := vertexCount * (vertexCount - 1) / 2
	if maxPossibleEdges == 0 {
		return 0
	}
	return float64(edgeCount) / float64(maxPossibleEdges)
}
