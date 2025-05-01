package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtreeCompare(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil || a.Val != b.Val {
		return false
	}

	return isSubtreeCompare(a.Right, b.Right) && isSubtreeCompare(a.Left, b.Left)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val && isSubtreeCompare(root, subRoot) {
		return true
	}

	return isSubtree(root.Right, subRoot) || isSubtree(root.Left, subRoot)
}
