package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	disjointSets := NewUnionFind(5)

	disjointSets.Union(0, 1)
	disjointSets.Union(2, 3)

	if !disjointSets.Connected(0, 1) {
		t.Error("0 and 1 should be connected")
	}
	if disjointSets.Connected(0, 2) {
		t.Error("0 and 2 should not be connected")
	}
	if disjointSets.Count() != 3 {
		t.Errorf("Expected 3 components, got %d", disjointSets.Count())
	}

	disjointSets.Union(1, 3)
	if !disjointSets.Connected(0, 3) {
		t.Error("0 and 3 should be connected after transitive union")
	}
	if disjointSets.Count() != 2 {
		t.Errorf("Expected 2 components, got %d", disjointSets.Count())
	}
}
