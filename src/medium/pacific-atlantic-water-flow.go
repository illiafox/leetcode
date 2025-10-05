package medium

// https://leetcode.com/problems/pacific-atlantic-water-flow/
func pacificAtlanticBFS(queue [][2]int, heights [][]int) [][]bool {
	visited := make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		visited[i] = make([]bool, len(heights[i]))
	}

	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]

		if visited[i][j] {
			continue
		}
		visited[i][j] = true

		for _, direction := range directions {
			newI, newJ := i+direction[0], j+direction[1]
			if newI < 0 || newI >= len(heights) || newJ < 0 || newJ >= len(heights[0]) {
				continue
			}

			if visited[newI][newJ] {
				continue
			}
			if heights[newI][newJ] >= heights[i][j] {
				queue = append(queue, [2]int{newI, newJ})
			}
		}
	}

	return visited
}

func pacificAtlantic(heights [][]int) (res [][]int) {
	n := len(heights) - 1
	m := len(heights[0]) - 1

	pacificQueue := make([][2]int, 0, n+m)
	atlanticQueue := make([][2]int, 0, n+m)

	for i := 0; i <= n; i++ {
		pacificQueue = append(pacificQueue, [2]int{i, 0})
		atlanticQueue = append(atlanticQueue, [2]int{i, m})
	}

	for j := 0; j <= m; j++ {
		pacificQueue = append(pacificQueue, [2]int{0, j})
		atlanticQueue = append(atlanticQueue, [2]int{n, j})
	}

	pacific := pacificAtlanticBFS(pacificQueue, heights)
	atlantic := pacificAtlanticBFS(atlanticQueue, heights)

	for i := 0; i < len(pacific); i++ {
		for j := 0; j < len(pacific[i]); j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
