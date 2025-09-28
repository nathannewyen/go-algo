package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	graphInstance := NewGraph(5, false)
	graphInstance.AddEdge(0, 1)
	graphInstance.AddEdge(0, 2)
	graphInstance.AddEdge(1, 3)
	graphInstance.AddEdge(2, 4)

	bfsResult := graphInstance.BFS(0)
	if len(bfsResult) != 5 {
		t.Errorf("BFS should visit all 5 vertices, got %d", len(bfsResult))
	}
	// First vertex visited should be the start vertex
	if bfsResult[0] != 0 {
		t.Errorf("BFS should start at vertex 0, got %d", bfsResult[0])
	}
}

func TestDFS(t *testing.T) {
	graphInstance := NewGraph(5, false)
	graphInstance.AddEdge(0, 1)
	graphInstance.AddEdge(0, 2)
	graphInstance.AddEdge(1, 3)
	graphInstance.AddEdge(2, 4)

	dfsResult := graphInstance.DFS(0)
	if len(dfsResult) != 5 {
		t.Errorf("DFS should visit all 5 vertices, got %d", len(dfsResult))
	}
	if dfsResult[0] != 0 {
		t.Errorf("DFS should start at vertex 0, got %d", dfsResult[0])
	}
}

func TestTopologicalSort(t *testing.T) {
	directedGraph := NewGraph(4, true)
	directedGraph.AddEdge(0, 1)
	directedGraph.AddEdge(0, 2)
	directedGraph.AddEdge(1, 3)
	directedGraph.AddEdge(2, 3)

	sortedResult := directedGraph.TopologicalSort()
	if len(sortedResult) != 4 {
		t.Errorf("TopologicalSort should return all 4 vertices, got %d", len(sortedResult))
	}
}
