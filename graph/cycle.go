package graph

// HasCycle detects if a directed graph contains a cycle using DFS coloring
// White(0)=unvisited, Gray(1)=in-progress, Black(2)=completed
func (g *AdjacencyListGraph) HasCycle() bool {
	colorState := make(map[int]int)
	for vertex := 0; vertex < g.vertices; vertex++ {
		colorState[vertex] = 0
	}

	for vertex := 0; vertex < g.vertices; vertex++ {
		if colorState[vertex] == 0 {
			if g.detectCycleDFS(vertex, colorState) {
				return true
			}
		}
	}
	return false
}

// detectCycleDFS uses three-color DFS to find back edges indicating cycles
func (g *AdjacencyListGraph) detectCycleDFS(vertex int, colorState map[int]int) bool {
	// Mark vertex as in-progress
	colorState[vertex] = 1

	for _, neighbor := range g.adjacency[vertex] {
		// Back edge found - cycle exists
		if colorState[neighbor] == 1 {
			return true
		}
		// Recurse on unvisited vertices
		if colorState[neighbor] == 0 {
			if g.detectCycleDFS(neighbor, colorState) {
				return true
			}
		}
	}

	// Mark vertex as completed
	colorState[vertex] = 2
	return false
}
