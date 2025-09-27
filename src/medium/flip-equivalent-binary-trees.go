package medium

// https://leetcode.com/problems/flip-equivalent-binary-trees
func flipEquiv(a, b *TreeNode) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil || a.Val != b.Val {
		return false
	}

	if flipEquiv(a.Left, b.Left) && flipEquiv(a.Right, b.Right) {
		return true
	}

	return flipEquiv(a.Left, b.Right) && flipEquiv(a.Right, b.Left)
}
