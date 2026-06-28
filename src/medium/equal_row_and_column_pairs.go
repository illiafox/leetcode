package medium

func equalPairs(grid [][]int) int {
	n := len(grid)

	hashedRows := make(map[int][]int)
	hashedColumns := make(map[int][]int)

	for i := range n {
		s := 0
		for j := range n {
			s += grid[i][j]
		}

		hashedRows[s] = append(hashedRows[s], i)
	}

	for i := range n {
		s := 0
		for j := range n {
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
				for i := range n {
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
