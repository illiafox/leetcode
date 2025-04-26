package easy

func invertTreeTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeTraverse(root.Right)
	invertTreeTraverse(root.Left)
}

func invertTree(root *TreeNode) *TreeNode {
	invertTreeTraverse(root)
	return root
}
