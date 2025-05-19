package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverseSumNumbers(out *int, node *TreeNode, accSum int) {
	accSum = accSum*10 + node.Val

	if node.Left == nil && node.Right == nil {
		*out += accSum
	}

	if node.Left != nil {
		traverseSumNumbers(out, node.Left, accSum)
	}
	if node.Right != nil {
		traverseSumNumbers(out, node.Right, accSum)
	}
}

func sumNumbers(root *TreeNode) int {
	total := 0
	traverseSumNumbers(&total, root, 0)
	return total
}
