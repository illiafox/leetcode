package medium

import "testing"

func TestMinimumArea_Table(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		// empty grid
		{grid: [][]int{}, want: 0},
		{grid: [][]int{{}}, want: 0},

		// single cell
		{grid: [][]int{{0}}, want: 0},
		{grid: [][]int{{1}}, want: 1},

		// line of 1s
		{grid: [][]int{{1, 1, 1}}, want: 3}, // height=1, width=3 → 3

		// two rows with spread 1s
		{grid: [][]int{
			{0, 1, 0},
			{1, 0, 1},
		}, want: 6}, // height=2, width=3 → 6

		// scattered 1s at corners
		{grid: [][]int{
			{1, 0, 0},
			{0, 0, 0},
			{0, 0, 1},
		}, want: 9}, // minRow=0, maxRow=2, minCol=0, maxCol=2 → 3x3

		// all ones 3x3
		{grid: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}, want: 9}, // height=3, width=3 → 9
	}

	for _, tc := range tests {
		got := minimumArea(tc.grid)
		if got != tc.want {
			t.Fatalf("minimumArea=%d, want %d; grid=%v", got, tc.want, tc.grid)
		}
	}
}
