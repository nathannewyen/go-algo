package matrix

// CalculateIslandPerimeter computes perimeter of island in grid
// 1 represents land, 0 represents water
func CalculateIslandPerimeter(landGrid [][]int) int {
	totalPerimeter := 0
	rowCount := len(landGrid)
	columnCount := len(landGrid[0])

	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
			if landGrid[rowIndex][columnIndex] == 1 {
				// Each land cell starts with 4 sides
				cellContribution := 4

				// Subtract shared edges with adjacent land cells
				if rowIndex > 0 && landGrid[rowIndex-1][columnIndex] == 1 {
					cellContribution -= 2
				}
				if columnIndex > 0 && landGrid[rowIndex][columnIndex-1] == 1 {
					cellContribution -= 2
				}
				totalPerimeter += cellContribution
			}
		}
	}
	return totalPerimeter
}
