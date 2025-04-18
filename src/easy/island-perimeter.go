package easy

func islandPerimeter(grid [][]int) int {
	var directions = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	perimeter := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				for _, direction := range directions {
					newI, newJ := i+direction[0], j+direction[1]
					if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[i]) || grid[newI][newJ] == 0 {
						perimeter++
					}
				}
			}
		}
	}

	return perimeter
}
