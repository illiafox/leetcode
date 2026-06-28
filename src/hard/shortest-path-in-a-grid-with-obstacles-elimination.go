package hard

// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
//
// You are given an m x n integer matrix grid where each cell is either 0 (empty) or 1 (obstacle).
// You can move up, down, left, or right from and to an empty cell in one step.
// Return the minimum number of steps to walk from the upper left corner (0, 0)
// to the lower right corner (m - 1, n - 1) given that you can eliminate at most k obstacles.
// If it is not possible to find such walk return -1.
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 40 (important, max 1600 cells)
// 1 <= k <= m * n
// grid[i][j] is either 0 or 1.
// grid[0][0] == grid[m - 1][n - 1] == 0
func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if stepsNeeded := m + n - 2; k >= stepsNeeded {
		return stepsNeeded
	}

	// Bits:
	// 10 ... 0  - eliminated (11 bits, max 2048)
	// 16 ... 11 - j (6 bits, max 64)
	// 22 ... 17 - i (6 bits, max 64)
	// 31 ... 23 - unused
	type state uint32

	pack := func(i, j int, elim int16) state {
		return state(uint32(i<<17) | uint32(j<<11) | uint32(elim))
	}
	unpack := func(st state) (i, j int, elim int16) {
		return int((st >> 17) & 0b111111), // 6 bits for i
			int((st >> 11) & 0b111111), // 6 bits for j
			int16(st & 0b11111111111) // 11 bits for elim
	}

	queue := []state{0}
	var swapQueue []state

	gridMinEliminated := make([]int16, m*n) // 0 == not set

	dirY := [4]int{0, 0, 1, -1}
	dirX := [4]int{1, -1, 0, 0}

	step := 0
	for len(queue) > 0 {
		for _, st := range queue {
			i, j, eliminated := unpack(st)

			for dirIdx := range 4 {
				newI, newJ := i+dirY[dirIdx], j+dirX[dirIdx]
				if newI < 0 || newI >= m || newJ < 0 || newJ >= n {
					continue
				}

				if newI == m-1 && newJ == n-1 {
					return step + 1
				}

				minEliminated := eliminated
				if grid[newI][newJ] == 1 {
					minEliminated++

					if minEliminated > int16(k) {
						continue
					}
				}

				idx := newI*n + newJ
				stored := gridMinEliminated[idx]
				cellMinEliminated := minEliminated + 1

				switch {
				case stored == 0, cellMinEliminated < stored:
					gridMinEliminated[idx] = cellMinEliminated
				case cellMinEliminated >= stored:
					continue
				}

				swapQueue = append(swapQueue, pack(newI, newJ, minEliminated))
			}
		}

		queue, swapQueue = swapQueue, queue[:0]
		step += 1
	}

	return -1
}
