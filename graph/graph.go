package graph

// AdjacencyListGraph represents a graph using adjacency lists
type AdjacencyListGraph struct {
	vertices  int
	directed  bool
	adjacency map[int][]int
}

// NewGraph creates a new graph with the given number of vertices
func NewGraph(vertices int, directed bool) *AdjacencyListGraph {
	return &AdjacencyListGraph{
		vertices:  vertices,
		directed:  directed,
		adjacency: make(map[int][]int),
	}
}

// AddEdge adds an edge between two vertices
func (g *AdjacencyListGraph) AddEdge(source int, destination int) {
	g.adjacency[source] = append(g.adjacency[source], destination)
	// For undirected graphs, add edge in both directions
	if !g.directed {
		g.adjacency[destination] = append(g.adjacency[destination], source)
	}
}

// GetNeighbors returns all neighbors of a vertex
func (g *AdjacencyListGraph) GetNeighbors(vertex int) []int {
	return g.adjacency[vertex]
}

// GetVertices returns the number of vertices
func (g *AdjacencyListGraph) GetVertices() int {
	return g.vertices
}
