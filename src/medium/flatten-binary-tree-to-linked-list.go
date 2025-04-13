package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flattenTraversal(root, prev *TreeNode) *TreeNode {
	if root == nil {
		return prev
	}

	prev.Right = root
	prev.Left = nil
	prev = root

	rootCopy := *root
	prev = flattenTraversal(rootCopy.Left, prev)
	return flattenTraversal(rootCopy.Right, prev)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	rootCopy := *root
	prev := flattenTraversal(rootCopy.Left, root)
	flattenTraversal(rootCopy.Right, prev)
}
