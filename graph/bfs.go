package graph

// BFS performs breadth-first search starting from the given vertex
func (g *AdjacencyListGraph) BFS(startVertex int) []int {
	visited := make(map[int]bool)
	traversalOrder := []int{}
	// Queue for BFS processing
	queue := []int{startVertex}
	visited[startVertex] = true

	for len(queue) > 0 {
		// Dequeue the front vertex
		currentVertex := queue[0]
		queue = queue[1:]
		traversalOrder = append(traversalOrder, currentVertex)

		// Visit all unvisited neighbors
		for _, neighbor := range g.adjacency[currentVertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return traversalOrder
}
