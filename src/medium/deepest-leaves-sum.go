package medium

func deepestLeavesMaxDepth(out *int, node *TreeNode, depth int) {
	if node == nil {
		return
	}

	depth++

	if node.Left == nil && node.Right == nil {
		*out = max(*out, depth)
		return
	}

	if node.Left != nil {
		deepestLeavesMaxDepth(out, node.Left, depth)
	}
	if node.Right != nil {
		deepestLeavesMaxDepth(out, node.Right, depth)
	}
}

func deepestLeavesSumTraverse(out *int, node *TreeNode, targetDepth int, depth int) {
	if node == nil {
		return
	}

	depth++

	if node.Left == nil && node.Right == nil {
		if depth == targetDepth {
			*out += node.Val
		}
		return
	}

	if node.Left != nil {
		deepestLeavesSumTraverse(out, node.Left, targetDepth, depth)
	}
	if node.Right != nil {
		deepestLeavesSumTraverse(out, node.Right, targetDepth, depth)
	}
}

// https://leetcode.com/problems/deepest-leaves-sum
func deepestLeavesSum(root *TreeNode) int {
	maxDepth := 0
	deepestLeavesMaxDepth(&maxDepth, root, 0)

	out := 0
	deepestLeavesSumTraverse(&out, root, maxDepth, 0)

	return out
}

func deepestLeavesSumBFS(root *TreeNode) int {
	deepest := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		sum := 0
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		deepest = sum
	}
	return deepest
}
