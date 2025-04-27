package easy

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	target := root.Val

	queue := []*TreeNode{root.Right, root.Left}

	for len(queue) > 0 {
		pop := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if pop == nil {
			continue
		}

		if pop.Val != target {
			return false
		}

		if pop.Right != nil {
			queue = append(queue, pop.Right)
		}

		if pop.Left != nil {
			queue = append(queue, pop.Left)
		}
	}

	return true
}

func isUnivalTreeRecursiveTraverse(root *TreeNode, target int) bool {
	if root == nil {
		return true
	}

	if root.Val != target {
		return false
	}

	return isUnivalTreeRecursiveTraverse(root.Left, target) && isUnivalTreeRecursiveTraverse(root.Right, target)
}

func isUnivalTreeRecursive(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isUnivalTreeRecursiveTraverse(root, root.Val)
}
