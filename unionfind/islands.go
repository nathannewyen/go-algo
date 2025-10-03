package unionfind

// CountIslands counts connected land regions in a grid using union-find
func CountIslands(landGrid [][]byte) int {
	if len(landGrid) == 0 {
		return 0
	}

	rowCount := len(landGrid)
	columnCount := len(landGrid[0])
	waterCellCount := 0
	islandUnionFind := NewUnionFind(rowCount * columnCount)

	for currentRow := 0; currentRow < rowCount; currentRow++ {
		for currentColumn := 0; currentColumn < columnCount; currentColumn++ {
			if landGrid[currentRow][currentColumn] == 0 {
				waterCellCount++
				continue
			}
			// Convert 2D coordinates to 1D index
			currentCellIndex := currentRow*columnCount + currentColumn

			// Union with right neighbor
			if currentColumn+1 < columnCount && landGrid[currentRow][currentColumn+1] == 1 {
				rightNeighborIndex := currentRow*columnCount + currentColumn + 1
				islandUnionFind.Union(currentCellIndex, rightNeighborIndex)
			}
			// Union with bottom neighbor
			if currentRow+1 < rowCount && landGrid[currentRow+1][currentColumn] == 1 {
				bottomNeighborIndex := (currentRow+1)*columnCount + currentColumn
				islandUnionFind.Union(currentCellIndex, bottomNeighborIndex)
			}
		}
	}
	return islandUnionFind.Count() - waterCellCount
}
