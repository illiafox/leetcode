package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// direction - true for right, false for left
func zigZagTraverse(root *TreeNode, pathCount int, direction bool) int {
	if root == nil {
		return pathCount
	}

	pathCount++

	right, left := pathCount, pathCount
	if direction {
		right = 0
	} else {
		left = 0
	}

	return max(
		zigZagTraverse(root.Right, right, !direction),
		zigZagTraverse(root.Left, left, !direction),
	)
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(
		zigZagTraverse(root, 0, true),
		zigZagTraverse(root, 0, false),
	) - 1
}
