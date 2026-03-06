package matrix

// Matrix represents a 2D integer matrix
type Matrix struct {
	rowCount    int
	columnCount int
	elements    [][]int
}

// NewMatrix creates a new matrix with given dimensions initialized to zero
func NewMatrix(rowCount int, columnCount int) *Matrix {
	elements := make([][]int, rowCount)
	for rowIndex := range elements {
		elements[rowIndex] = make([]int, columnCount)
	}
	return &Matrix{rowCount: rowCount, columnCount: columnCount, elements: elements}
}

// Get returns the element at the specified row and column
func (matrix *Matrix) Get(rowIndex int, columnIndex int) int {
	return matrix.elements[rowIndex][columnIndex]
}

// Set assigns a value to the specified row and column
func (matrix *Matrix) Set(rowIndex int, columnIndex int, value int) {
	matrix.elements[rowIndex][columnIndex] = value
}
