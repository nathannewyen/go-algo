package graph

// ShortestPathBFS finds shortest path between two vertices in unweighted graph
// Returns the path as a slice of vertices, empty if no path exists
func (g *AdjacencyListGraph) ShortestPathBFS(startVertex int, endVertex int) []int {
	visited := make(map[int]bool)
	// Track parent of each vertex to reconstruct path
	parentVertex := make(map[int]int)
	parentVertex[startVertex] = -1
	visited[startVertex] = true
	queue := []int{startVertex}

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]

		if currentVertex == endVertex {
			return reconstructPath(parentVertex, endVertex)
		}

		for _, neighbor := range g.adjacency[currentVertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parentVertex[neighbor] = currentVertex
				queue = append(queue, neighbor)
			}
		}
	}
	return []int{}
}

// reconstructPath builds the path from start to end using parent pointers
func reconstructPath(parentVertex map[int]int, endVertex int) []int {
	path := []int{}
	for currentVertex := endVertex; currentVertex != -1; currentVertex = parentVertex[currentVertex] {
		path = append([]int{currentVertex}, path...)
	}
	return path
}
