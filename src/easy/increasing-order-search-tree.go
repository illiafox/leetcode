package easy

func increasingBSTTraverse(root *TreeNode, curr **TreeNode) {
	if root == nil {
		return
	}

	increasingBSTTraverse(root.Left, curr)

	root.Left = nil

	(*curr).Right = root
	*curr = root

	increasingBSTTraverse(root.Right, curr)
}

// https://leetcode.com/problems/increasing-order-search-tree/
func increasingBST(root *TreeNode) *TreeNode {
	stub := &TreeNode{}
	curr := stub
	increasingBSTTraverse(root, &curr)

	return stub.Right
}
