package hard

import (
	"math"
	"slices"
)

// https://www.geeksforgeeks.org/0-1-bfs-algorithm-for-competitive-programming/

func minCost(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	minCosts := make([][]int, m)
	for i := range minCosts {
		minCosts[i] = make([]int, n)
		for j := range minCosts[i] {
			minCosts[i][j] = math.MaxInt32
		}
	}
	minCosts[0][0] = 0

	moves := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	queue := [][2]int{{0, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		j, i := current[1], current[0]

		for moveIdx, move := range moves {
			newJ, newI := j+move[0], i+move[1]
			if newJ < 0 || newJ >= n || newI < 0 || newI >= m {
				continue
			}

			cost := 1
			if moveIdx+1 == grid[i][j] {
				cost = 0
			}

			distance := minCosts[i][j] + cost

			if distance < minCosts[newI][newJ] {
				minCosts[newI][newJ] = distance

				coords := [2]int{newI, newJ}
				if cost == 1 {
					queue = append(queue, coords)
				} else {
					queue = slices.Insert(queue, 0, coords)
				}
			}
		}
	}

	return minCosts[m-1][n-1]
}
