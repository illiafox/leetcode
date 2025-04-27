package easy

func inorderTraversal(root *TreeNode) (out []int) {
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		out = append(out, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return out
}
