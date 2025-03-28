package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverseDepth(root *TreeNode, n int) int {
	if root == nil {
		return n
	}

	return max(traverseDepth(root.Left, n+1), traverseDepth(root.Right, n+1))
}

func maxDepth(root *TreeNode) int {
	return traverseDepth(root, 0)
}
