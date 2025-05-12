package easy

// solution for https://leetcode.com/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0

	var traverse func(*TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := traverse(node.Left)
		right := traverse(node.Right)

		maxLength = max(maxLength, left+right)

		return 1 + max(left, right)
	}

	traverse(root)
	return maxLength
}
