package medium

func addOneRowTraverse(root *TreeNode, prev **TreeNode, val int, targetDepth int, curDepth int, leftDirection bool) {
	if curDepth++; curDepth == targetDepth {
		newNode := &TreeNode{Val: val}
		if leftDirection {
			newNode.Left = root
		} else {
			newNode.Right = root
		}
		*prev = newNode
	}

	if root == nil {
		return
	}

	addOneRowTraverse(root.Right, &root.Right, val, targetDepth, curDepth, false)
	addOneRowTraverse(root.Left, &root.Left, val, targetDepth, curDepth, true)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	out := root
	addOneRowTraverse(root, &out, val, depth, 0, true)
	return out
}
