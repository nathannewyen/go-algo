package unionfind

// FindRedundantConnection finds the edge that creates a cycle in the graph
// Returns the last such edge encountered
func FindRedundantConnection(edgeList [][]int) []int {
	// Find maximum vertex number to size the union-find
	maxVertex := 0
	for _, edge := range edgeList {
		if edge[0] > maxVertex {
			maxVertex = edge[0]
		}
		if edge[1] > maxVertex {
			maxVertex = edge[1]
		}
	}

	cycleDetector := NewUnionFind(maxVertex + 1)

	for _, edge := range edgeList {
		// If both vertices already connected, this edge creates a cycle
		if cycleDetector.Connected(edge[0], edge[1]) {
			return edge
		}
		cycleDetector.Union(edge[0], edge[1])
	}
	return []int{}
}
