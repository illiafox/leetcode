package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func removeLeafNodesTraverse(node *TreeNode, prev **TreeNode, target int) {
	if node == nil {
		return
	}

	removeLeafNodesTraverse(node.Left, &node.Left, target)
	removeLeafNodesTraverse(node.Right, &node.Right, target)

	if node.Left == nil && node.Right == nil && node.Val == target {
		*prev = nil
	}
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	removeLeafNodesTraverse(root, &root, target)

	return root
}
