package medium

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	row := 0
	col := cols - 1

	for row < rows && col >= 0 {
		current := matrix[row][col]
		if current == target {
			return true
		} else if current > target {
			col-- // Move left
		} else {
			row++ // Move down
		}
	}

	return false
}
