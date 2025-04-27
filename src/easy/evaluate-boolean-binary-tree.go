package easy

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	switch root.Val {
	case 2:
		return evaluateTree(root.Right) || evaluateTree(root.Left)
	case 3:
		return evaluateTree(root.Right) && evaluateTree(root.Left)
	}

	return root.Val == 1
}
