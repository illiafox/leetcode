package easy

func hasPathSumTraverse(root *TreeNode, cur, targetSum int) bool {
	if root == nil {
		return false
	}

	cur += root.Val
	if cur == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSumTraverse(root.Left, cur, targetSum) || hasPathSumTraverse(root.Right, cur, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumTraverse(root, 0, targetSum)
}
