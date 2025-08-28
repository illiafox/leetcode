package medium

import "sort"

// https://leetcode.com/problems/sort-matrix-by-diagonals
// the grid is guaranteed n×n
func sortMatrix(grid [][]int) [][]int {
	// bottom-left triangle
	for y := 0; y < len(grid); y++ {
		sort.Sort(diagonalSort{
			grid:   grid,
			startY: y,
			startX: 0,
			desc:   false,
		})
	}

	// top-right triangle
	for y := 1; y < len(grid)-1; y++ {
		sort.Sort(diagonalSort{
			grid:   grid,
			startY: 0,
			startX: y,
			desc:   true,
		})
	}

	return grid
}

type diagonalSort struct {
	grid   [][]int
	startX int
	startY int
	desc   bool
}

func (d diagonalSort) Len() int {
	return len(d.grid) - d.startY - d.startX
}

func (d diagonalSort) Less(i, j int) bool {
	iY, iX := d.startY+i, i+d.startX
	jY, jX := d.startY+j, j+d.startX

	if d.desc {
		return d.grid[iY][iX] < d.grid[jY][jX]
	}

	return d.grid[iY][iX] > d.grid[jY][jX]
}

func (d diagonalSort) Swap(i, j int) {
	iY, iX := d.startY+i, i+d.startX
	jY, jX := d.startY+j, j+d.startX

	d.grid[iY][iX], d.grid[jY][jX] = d.grid[jY][jX], d.grid[iY][iX]
}
