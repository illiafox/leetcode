package medium

// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i
func minimumArea(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	minCol, maxCol := len(grid[0]), 0
	minRow, maxRow := len(grid), 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				minCol = min(minCol, j)
				minRow = min(minRow, i)
				maxCol = max(maxCol, j)
				maxRow = max(maxRow, i)
			}
		}
	}

	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}
