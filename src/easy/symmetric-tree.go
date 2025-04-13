package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetricTraversal(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil) != (b == nil) {
		return false
	}

	return a.Val == b.Val &&
		isSymmetricTraversal(a.Left, b.Right) &&
		isSymmetricTraversal(a.Right, b.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTraversal(root.Left, root.Right)
}
