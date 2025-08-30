package medium

// https://leetcode.com/problems/valid-sudoku
func isValidSudoku(board [][]byte) bool {
	var rows, cols, blocks [9]uint16

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			elem := board[i][j]
			if elem == '.' {
				continue
			}
			if elem < '1' || elem > '9' {
				return false
			}

			bit := uint16(1) << (elem - '1')

			blockIdx := (i/3)*3 + (j / 3)

			if rows[i]&bit != 0 || cols[j]&bit != 0 || blocks[blockIdx]&bit != 0 {
				return false
			}

			rows[i] |= bit
			cols[j] |= bit
			blocks[blockIdx] |= bit
		}
	}

	return true
}

func isValidSudokuOld(board [][]byte) bool {
	set := make(map[byte]struct{}, 9)

	check := func(elem byte) bool {
		if elem == '.' {
			return false
		}
		if elem < '1' || elem > '9' {
			return true
		}
		if _, ok := set[elem]; ok {
			return true
		}
		set[elem] = struct{}{}
		return false
	}

	for _, row := range board {
		for _, elem := range row {
			if check(elem) {
				return false
			}
		}
		clear(set)
	}

	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			if check(board[i][j]) {
				return false
			}
		}
		clear(set)
	}

	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			for j := (x - 1) * 3; j < x*3; j++ {
				for i := (y - 1) * 3; i < y*3; i++ {
					if check(board[i][j]) {
						return false
					}
				}
			}
			clear(set)
		}
	}

	return true
}
