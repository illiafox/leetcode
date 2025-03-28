package easy

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLeafs(root *TreeNode, in []int) []int {
	if root == nil {
		return in
	}

	if root.Left == nil && root.Right == nil {
		in = append(in, root.Val)
		return in
	}

	in = getLeafs(root.Left, in)
	in = getLeafs(root.Right, in)

	return in
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return slices.Equal(getLeafs(root1, nil), getLeafs(root2, nil))
}
