package medium

func averageOfSubtree(root *TreeNode) int {
	var count int
	averageOfSubtreePostOrder(root, &count)
	return count
}

func averageOfSubtreePostOrder(node *TreeNode, count *int) (sum, nodes int) {
	if node == nil {
		return 0, 0
	}

	leftSum, leftNodes := averageOfSubtreePostOrder(node.Left, count)
	rightSum, rightNodes := averageOfSubtreePostOrder(node.Right, count)

	sum = node.Val + leftSum + rightSum
	nodes = 1 + leftNodes + rightNodes

	if sum/nodes == node.Val {
		*count++
	}

	return sum, nodes
}
