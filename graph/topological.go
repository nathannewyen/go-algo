package graph

// TopologicalSort returns vertices in topological order using Kahns algorithm
// Only valid for directed acyclic graphs (DAGs)
func (g *AdjacencyListGraph) TopologicalSort() []int {
	// Calculate in-degree for each vertex
	inDegree := make(map[int]int)
	for i := 0; i < g.vertices; i++ {
		inDegree[i] = 0
	}
	for _, neighbors := range g.adjacency {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	// Start with vertices that have no incoming edges
	queue := []int{}
	for vertex := 0; vertex < g.vertices; vertex++ {
		if inDegree[vertex] == 0 {
			queue = append(queue, vertex)
		}
	}

	sortedOrder := []int{}
	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, currentVertex)

		// Reduce in-degree for all neighbors
		for _, neighbor := range g.adjacency[currentVertex] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return sortedOrder
}
