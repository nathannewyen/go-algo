package unionfind

import "sort"

// WeightedEdge represents an edge with a weight connecting two vertices
type WeightedEdge struct {
	SourceVertex      int
	DestinationVertex int
	EdgeWeight        int
}

// KruskalMST finds minimum spanning tree using Kruskals algorithm
// Returns edges in the MST and total weight
func KruskalMST(vertexCount int, edgeList []WeightedEdge) ([]WeightedEdge, int) {
	// Sort edges by weight in ascending order (greedy choice)
	sort.Slice(edgeList, func(i int, j int) bool {
		return edgeList[i].EdgeWeight < edgeList[j].EdgeWeight
	})

	mstUnionFind := NewUnionFind(vertexCount)
	mstEdges := []WeightedEdge{}
	totalMSTWeight := 0

	for _, candidateEdge := range edgeList {
		// Only add edge if it connects two different components (no cycle)
		if !mstUnionFind.Connected(candidateEdge.SourceVertex, candidateEdge.DestinationVertex) {
			mstUnionFind.Union(candidateEdge.SourceVertex, candidateEdge.DestinationVertex)
			mstEdges = append(mstEdges, candidateEdge)
			totalMSTWeight += candidateEdge.EdgeWeight
		}
		// MST is complete when we have V-1 edges
		if len(mstEdges) == vertexCount-1 {
			break
		}
	}
	return mstEdges, totalMSTWeight
}
