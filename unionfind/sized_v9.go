package unionfind

// SizedUnionFind9 tracks component sizes alongside union operations
type SizedUnionFind9 struct {
	parentNode     []int
	componentSize  []int
	totalComponents int
}

// NewSizedUnionFind9 creates a size-tracking union-find
func NewSizedUnionFind9(elementCount int) *SizedUnionFind9 {
	parentNode := make([]int, elementCount)
	componentSize := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		parentNode[i] = i
		componentSize[i] = 1
	}
	return &SizedUnionFind9{parentNode: parentNode, componentSize: componentSize, totalComponents: elementCount}
}

// GetComponentSize9 returns the size of the component containing element
func (unionFind *SizedUnionFind9) GetComponentSize9(element int) int {
	rootElement := element
	for unionFind.parentNode[rootElement] != rootElement {
		rootElement = unionFind.parentNode[rootElement]
	}
	return unionFind.componentSize[rootElement]
}
