package unionfind

// UnionFind implements disjoint set data structure with path compression
type UnionFind struct {
	parentNode    []int
	componentRank []int
	componentCount int
}

// NewUnionFind creates a union-find structure where each element is its own set
func NewUnionFind(elementCount int) *UnionFind {
	parentNode := make([]int, elementCount)
	componentRank := make([]int, elementCount)
	for elementIndex := 0; elementIndex < elementCount; elementIndex++ {
		parentNode[elementIndex] = elementIndex
	}
	return &UnionFind{
		parentNode:    parentNode,
		componentRank: componentRank,
		componentCount: elementCount,
	}
}

// Find returns the root representative of the set containing element
// Uses path compression to flatten the tree structure
func (unionFind *UnionFind) Find(element int) int {
	if unionFind.parentNode[element] != element {
		// Path compression: point directly to root
		unionFind.parentNode[element] = unionFind.Find(unionFind.parentNode[element])
	}
	return unionFind.parentNode[element]
}

// Union merges the sets containing elementA and elementB
// Uses union by rank to keep tree balanced
func (unionFind *UnionFind) Union(elementA int, elementB int) {
	rootA := unionFind.Find(elementA)
	rootB := unionFind.Find(elementB)

	if rootA == rootB {
		return
	}

	// Attach smaller rank tree under root of higher rank tree
	if unionFind.componentRank[rootA] < unionFind.componentRank[rootB] {
		unionFind.parentNode[rootA] = rootB
	} else if unionFind.componentRank[rootA] > unionFind.componentRank[rootB] {
		unionFind.parentNode[rootB] = rootA
	} else {
		unionFind.parentNode[rootB] = rootA
		unionFind.componentRank[rootA]++
	}
	unionFind.componentCount--
}

// Connected checks if two elements belong to the same set
func (unionFind *UnionFind) Connected(elementA int, elementB int) bool {
	return unionFind.Find(elementA) == unionFind.Find(elementB)
}

// Count returns the number of disjoint sets
func (unionFind *UnionFind) Count() int {
	return unionFind.componentCount
}
