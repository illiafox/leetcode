package medium

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := lcaDeepestLeavesDFS(root)
	return lca
}

func lcaDeepestLeavesDFS(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	lDepth, l := lcaDeepestLeavesDFS(root.Left)
	rDepth, r := lcaDeepestLeavesDFS(root.Right)

	if lDepth > rDepth {
		return lDepth + 1, l
	}
	if lDepth < rDepth {
		return rDepth + 1, r
	}
	return lDepth + 1, root
}
