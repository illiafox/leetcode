package medium

func equalPairs(grid [][]int) int {
	n := len(grid)

	hashedRows := make(map[int][]int)
	hashedColumns := make(map[int][]int)

	for i := 0; i < n; i++ {
		s := 0
		for j := 0; j < n; j++ {
			s += grid[i][j]
		}

		hashedRows[s] = append(hashedRows[s], i)
	}

	for i := 0; i < n; i++ {
		s := 0
		for j := 0; j < n; j++ {
			s += grid[j][i]
		}

		hashedColumns[s] = append(hashedColumns[s], i)
	}

	totalPairs := 0

	for rowHash, v := range hashedRows {
		cols := hashedColumns[rowHash]
		for _, rowIdx := range v {
			for _, colIdx := range cols {
				equal := true
				for i := 0; i < n; i++ {
					if grid[rowIdx][i] != grid[i][colIdx] {
						equal = false
						break
					}
				}

				if equal {
					totalPairs++
				}
			}
		}

	}

	return totalPairs
}
