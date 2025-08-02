package hard

import (
	"math"
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
type triple struct{ col, row, val int }

// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	triples := make([]triple, 0)

	var dfs func(n *TreeNode, col, row int)
	dfs = func(n *TreeNode, col, row int) {
		if n == nil {
			return
		}
		triples = append(triples, triple{col, row, n.Val})
		dfs(n.Left, col-1, row+1)
		dfs(n.Right, col+1, row+1)
	}
	dfs(root, 0, 0)

	sort.Slice(triples, func(i, j int) bool {
		if triples[i].col != triples[j].col {
			return triples[i].col < triples[j].col
		}
		if triples[i].row != triples[j].row {
			return triples[i].row < triples[j].row
		}
		return triples[i].val < triples[j].val
	})

	res := make([][]int, 0)
	curCol := math.MinInt
	for _, t := range triples {
		if t.col != curCol {
			res = append(res, []int{})
			curCol = t.col
		}
		res[len(res)-1] = append(res[len(res)-1], t.val)
	}
	return res
}
