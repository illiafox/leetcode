package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSudoku(t *testing.T) {
	type testCase struct {
		name     string
		board    [][]byte
		expected bool
	}

	mk := func(rows []string) [][]byte {
		b := make([][]byte, len(rows))
		for i := range rows {
			b[i] = []byte(rows[i])
		}
		return b
	}

	cases := []testCase{
		{"valid board", mk([]string{
			"53..7....",
			"6..195...",
			".98....6.",
			"8...6...3",
			"4..8.3..1",
			"7...2...6",
			".6....28.",
			"...419..5",
			"....8..79",
		}), true},
		{"duplicate in row", mk([]string{
			"53..7..5.", // duplicate '5' in row
			"6..195...",
			".98....6.",
			"8...6...3",
			"4..8.3..1",
			"7...2...6",
			".6....28.",
			"...419..5",
			"....8..79",
		}), false},
		{"duplicate in column", mk([]string{
			"53..7....",
			"6..195...",
			"6.8....6.", // duplicate '6' in col 0
			"8...6...3",
			"4..8.3..1",
			"7...2...6",
			".6....28.",
			"...419..5",
			"....8..79",
		}), false},
		{"duplicate in block", mk([]string{
			"53..7....",
			"3..195...", // duplicate '3' in 3x3 block
			".98....6.",
			"8...6...3",
			"4..8.3..1",
			"7...2...6",
			".6....28.",
			"...419..5",
			"....8..79",
		}), false},
		{"invalid character", mk([]string{
			"53..7....",
			"6..195...",
			".98....6.",
			"8...6...3",
			"4..8.3..1",
			"7...2...6",
			".6....28.",
			"...419..5",
			"..x.8..79", // invalid rune
		}), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, isValidSudoku(c.board))
			assert.Equal(t, c.expected, isValidSudokuOld(c.board))
		})
	}
}
