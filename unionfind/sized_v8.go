package unionfind

// SizedUnionFind8 tracks component sizes alongside union operations
type SizedUnionFind8 struct {
	parentNode     []int
	componentSize  []int
	totalComponents int
}

// NewSizedUnionFind8 creates a size-tracking union-find
func NewSizedUnionFind8(elementCount int) *SizedUnionFind8 {
	parentNode := make([]int, elementCount)
	componentSize := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		parentNode[i] = i
		componentSize[i] = 1
	}
	return &SizedUnionFind8{parentNode: parentNode, componentSize: componentSize, totalComponents: elementCount}
}

// GetComponentSize8 returns the size of the component containing element
func (unionFind *SizedUnionFind8) GetComponentSize8(element int) int {
	rootElement := element
	for unionFind.parentNode[rootElement] != rootElement {
		rootElement = unionFind.parentNode[rootElement]
	}
	return unionFind.componentSize[rootElement]
}
