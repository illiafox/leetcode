package hard

import (
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type position struct {
	val int
	row int
}

func verticalTraverse(cols map[int][]position, node *TreeNode, col, row int) {
	if node == nil {
		return
	}

	cols[col] = append(cols[col], position{row: row, val: node.Val})
	row++
	verticalTraverse(cols, node.Left, col-1, row)
	verticalTraverse(cols, node.Right, col+1, row)
}

// not final solution, see verticalTraversal
func verticalTraversalFirstTry(root *TreeNode) [][]int {
	cols := make(map[int][]position)
	if root == nil {
		return nil
	}

	verticalTraverse(cols, root, 0, 0)

	colPositions := make([][]position, len(cols))
	indexes := make([]int, len(cols))

	i := 0
	for col, v := range cols {
		indexes[i] = col
		colPositions[i] = v
		i++
	}

	sort.Sort(sortable{
		indexes: indexes,
		out:     colPositions,
	})

	out := make([][]int, len(cols))
	for j, positions := range colPositions {
		sort.Slice(positions, func(i, j int) bool {
			if positions[i].row == positions[j].row {
				return positions[i].val < positions[j].val
			}

			return positions[i].row < positions[j].row
		})

		out[j] = make([]int, len(positions))
		for k, p := range positions {
			out[j][k] = p.val
		}
	}

	return out
}

type sortable struct {
	out     [][]position
	indexes []int
}

func (s sortable) Len() int {
	return len(s.out)
}

func (s sortable) Less(i, j int) bool {
	return s.indexes[i] < s.indexes[j]
}

func (s sortable) Swap(i, j int) {
	s.indexes[i], s.indexes[j] = s.indexes[j], s.indexes[i]
	s.out[i], s.out[j] = s.out[j], s.out[i]
}
