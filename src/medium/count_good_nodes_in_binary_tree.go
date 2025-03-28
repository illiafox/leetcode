package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkNode(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}

	total := 0
	if root.Val >= x {
		total++
	}

	if root.Val > x {
		x = root.Val
	}

	return total + checkNode(root.Left, x) + checkNode(root.Right, x)
}

func goodNodes(root *TreeNode) int {
	return 1 + checkNode(root.Left, root.Val) + checkNode(root.Right, root.Val)
}
