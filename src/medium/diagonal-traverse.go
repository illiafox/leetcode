package medium

// https://leetcode.com/problems/diagonal-traverse
func findDiagonalOrder(mat [][]int) []int {
	rows := len(mat)
	cols := len(mat[0])

	out := make([]int, 0, rows*cols)
	temp := make([]int, 0, cols)

	down := false
	for d := 0; d <= rows+cols-2; d++ {
		r := max(0, d-(cols-1))
		c := min(d, cols-1)

		temp = temp[:0]
		for r < rows && c >= 0 {
			temp = append(temp, mat[r][c])
			r++
			c--
		}

		if !down {
			for i := len(temp) - 1; i >= 0; i-- {
				out = append(out, temp[i])
			}
		} else {
			out = append(out, temp...)
		}

		down = !down
	}

	return out
}
