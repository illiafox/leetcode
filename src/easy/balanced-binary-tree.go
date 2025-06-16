package easy

func isBalancedTraverse(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, l := isBalancedTraverse(root.Left)
	rightDepth, r := isBalancedTraverse(root.Right)
	if !l || !r || max(leftDepth, rightDepth)-min(leftDepth, rightDepth) > 1 {
		return 0, false
	}

	return max(leftDepth, rightDepth) + 1, true
}

func isBalanced(root *TreeNode) bool {
	_, ok := isBalancedTraverse(root)
	return ok
}
