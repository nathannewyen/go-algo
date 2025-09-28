package graph

// DFS performs depth-first search starting from the given vertex
func (g *AdjacencyListGraph) DFS(startVertex int) []int {
	visited := make(map[int]bool)
	traversalOrder := []int{}
	g.dfsRecursive(startVertex, visited, &traversalOrder)
	return traversalOrder
}

// dfsRecursive is the recursive helper for DFS
func (g *AdjacencyListGraph) dfsRecursive(vertex int, visited map[int]bool, order *[]int) {
	visited[vertex] = true
	*order = append(*order, vertex)

	// Recursively visit all unvisited neighbors
	for _, neighbor := range g.adjacency[vertex] {
		if !visited[neighbor] {
			g.dfsRecursive(neighbor, visited, order)
		}
	}
}

// DFSIterative performs depth-first search using a stack instead of recursion
func (g *AdjacencyListGraph) DFSIterative(startVertex int) []int {
	visited := make(map[int]bool)
	traversalOrder := []int{}
	// Stack for iterative DFS
	stack := []int{startVertex}

	for len(stack) > 0 {
		// Pop from stack
		currentVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[currentVertex] {
			continue
		}
		visited[currentVertex] = true
		traversalOrder = append(traversalOrder, currentVertex)

		// Push unvisited neighbors onto stack
		for _, neighbor := range g.adjacency[currentVertex] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
	return traversalOrder
}
