package easy

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var traverse func(root *TreeNode, curDepth int) int
	traverse = func(root *TreeNode, curDepth int) int {
		if root == nil {
			return math.MaxInt
		}
		if root.Right == nil && root.Left == nil {
			return curDepth
		}
		curDepth++

		return min(
			traverse(root.Right, curDepth),
			traverse(root.Left, curDepth),
		)
	}

	return traverse(root, 1)
}
