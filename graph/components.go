package graph

// ConnectedComponents returns groups of connected vertices in an undirected graph
func (g *AdjacencyListGraph) ConnectedComponents() [][]int {
	visited := make(map[int]bool)
	allComponents := [][]int{}

	for vertex := 0; vertex < g.vertices; vertex++ {
		if !visited[vertex] {
			// Collect all vertices in this component
			currentComponent := []int{}
			g.collectComponent(vertex, visited, &currentComponent)
			allComponents = append(allComponents, currentComponent)
		}
	}
	return allComponents
}

// collectComponent gathers all vertices reachable from the given vertex
func (g *AdjacencyListGraph) collectComponent(vertex int, visited map[int]bool, component *[]int) {
	visited[vertex] = true
	*component = append(*component, vertex)

	for _, neighbor := range g.adjacency[vertex] {
		if !visited[neighbor] {
			g.collectComponent(neighbor, visited, component)
		}
	}
}
