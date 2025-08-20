package medium

import "testing"

func TestCountSquares_Table(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		// empty
		{matrix: [][]int{{}}, want: 0},

		// single cell
		{matrix: [][]int{{0}}, want: 0},
		{matrix: [][]int{{1}}, want: 1},

		// 1x3 line
		{matrix: [][]int{{1, 1, 1}}, want: 3},

		// full 2x2 => four 1x1 + one 2x2 = 5
		{matrix: [][]int{
			{1, 1},
			{1, 1},
		}, want: 5},

		// 3x3 with a hole in center => only 1x1 (8 ones)
		{matrix: [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, want: 8},

		// mixed rectangle
		// 1x1: 9, 2x2: 2 (at (0,0) and (1,1)), total = 11
		{matrix: [][]int{
			{1, 1, 0, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 0},
		}, want: 11},

		// all ones 3x3 => 1x1:9, 2x2:4, 3x3:1 => 14
		{matrix: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}, want: 14},
	}

	for _, tc := range tests {
		got := countSquares(tc.matrix)
		if got != tc.want {
			t.Fatalf("countSquares=%d, want %d; matrix=%v", got, tc.want, tc.matrix)
		}

		got = countSquaresDP(tc.matrix)
		if got != tc.want {
			t.Fatalf("countSquaresDP=%d, want %d; matrix=%v", got, tc.want, tc.matrix)
		}
	}
}
