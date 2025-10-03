package unionfind

// SizedUnionFind7 tracks component sizes alongside union operations
type SizedUnionFind7 struct {
	parentNode     []int
	componentSize  []int
	totalComponents int
}

// NewSizedUnionFind7 creates a size-tracking union-find
func NewSizedUnionFind7(elementCount int) *SizedUnionFind7 {
	parentNode := make([]int, elementCount)
	componentSize := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		parentNode[i] = i
		componentSize[i] = 1
	}
	return &SizedUnionFind7{parentNode: parentNode, componentSize: componentSize, totalComponents: elementCount}
}

// GetComponentSize7 returns the size of the component containing element
func (unionFind *SizedUnionFind7) GetComponentSize7(element int) int {
	rootElement := element
	for unionFind.parentNode[rootElement] != rootElement {
		rootElement = unionFind.parentNode[rootElement]
	}
	return unionFind.componentSize[rootElement]
}
