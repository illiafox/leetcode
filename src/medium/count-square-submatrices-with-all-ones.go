package medium

func countSquaresDP(a [][]int) int {
	if len(a) == 0 || len(a[0]) == 0 {
		return 0
	}
	m, n := len(a), len(a[0])

	dp := make([]int, n)
	total := 0

	for i := 0; i < m; i++ {
		prevDiag := 0
		for j := 0; j < n; j++ {
			up := dp[j]

			if a[i][j] == 0 {
				dp[j] = 0
				prevDiag = up
				continue
			}

			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				left := dp[j-1]
				diag := prevDiag

				dp[j] = 1 + min(left, up, diag)
			}

			total += dp[j]
			prevDiag = up
		}
	}
	return total
}

func isSquareOnes(a [][]int, i, j, side int) bool {
	for ni := i; ni <= i+side; ni++ {
		for nj := j; nj <= j+side; nj++ {
			if a[ni][nj] == 0 {
				return false
			}
		}
	}
	return true
}

func countSquares(matrix [][]int) int {
	count := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				for side := 0; side <= len(matrix); side++ {
					if i+side >= len(matrix) || j+side >= len(matrix[i]) {
						break
					}

					if !isSquareOnes(matrix, i, j, side) {
						break
					}

					count++
				}
			}
		}
	}

	return count
}
