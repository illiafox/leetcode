package medium

func insertIntoBSTTraverse(node *TreeNode, ptr **TreeNode, val int) {
	if node == nil {
		*ptr = &TreeNode{Val: val}
		return
	}

	if node.Val > val {
		insertIntoBSTTraverse(node.Left, &node.Left, val)
	} else {
		insertIntoBSTTraverse(node.Right, &node.Right, val)
	}
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	insertIntoBSTTraverse(root, &root, val)
	return root
}
