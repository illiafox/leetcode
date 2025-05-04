package medium

func exist(board [][]byte, word string) bool {
	var directions = [][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var dfs func(i, j, idx int) bool
	dfs = func(i, j, idx int) bool {
		if idx == len(word) {
			return true
		}

		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < rows && nj >= 0 && nj < cols &&
				!visited[ni][nj] && board[ni][nj] == word[idx] {

				visited[ni][nj] = true
				if dfs(ni, nj, idx+1) {
					return true
				}
				visited[ni][nj] = false
			}
		}

		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				visited[i][j] = true
				if dfs(i, j, 1) {
					return true
				}
				visited[i][j] = false
			}
		}
	}

	return false
}
