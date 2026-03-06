package matrix

// MultiplyMatrices computes the product of two matrices
// First matrix dimensions: rowsA x colsA, Second: rowsB x colsB
// Requires colsA == rowsB
func MultiplyMatrices(matrixA [][]int, matrixB [][]int) [][]int {
	rowCountA := len(matrixA)
	columnCountA := len(matrixA[0])
	columnCountB := len(matrixB[0])

	// Initialize result matrix with zeros
	productMatrix := make([][]int, rowCountA)
	for rowIndex := range productMatrix {
		productMatrix[rowIndex] = make([]int, columnCountB)
	}

	// Compute each element as dot product of row and column
	for rowIndex := 0; rowIndex < rowCountA; rowIndex++ {
		for columnIndex := 0; columnIndex < columnCountB; columnIndex++ {
			dotProductSum := 0
			for innerIndex := 0; innerIndex < columnCountA; innerIndex++ {
				dotProductSum += matrixA[rowIndex][innerIndex] * matrixB[innerIndex][columnIndex]
			}
			productMatrix[rowIndex][columnIndex] = dotProductSum
		}
	}
	return productMatrix
}
