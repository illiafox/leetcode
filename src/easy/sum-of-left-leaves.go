package easy

func sumOfLeftLeavesTraverse(root *TreeNode, left bool, res *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if left {
			*res += root.Val
		}
		return
	}

	sumOfLeftLeavesTraverse(root.Left, true, res)
	sumOfLeftLeavesTraverse(root.Right, false, res)
}

func sumOfLeftLeaves(root *TreeNode) (res int) {
	sumOfLeftLeavesTraverse(root, false, &res)
	return
}
