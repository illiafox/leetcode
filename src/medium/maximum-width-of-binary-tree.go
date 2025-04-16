package medium

func widthOfBinaryTree(root *TreeNode) int {
	type task struct {
		node  *TreeNode
		index int
	}

	var queue []task
	maxWidth := 0

	queue = append(queue, task{root, 0})
	for len(queue) > 0 {
		levelSize := len(queue)

		firstIndex := queue[0].index
		lastIndex := queue[levelSize-1].index

		maxWidth = max(maxWidth, lastIndex-firstIndex+1)

		for i := 0; i < levelSize; i++ {
			node := queue[i].node
			index := queue[i].index

			if node.Left != nil {
				queue = append(queue, task{node.Left, 2 * index})
			}

			if node.Right != nil {
				queue = append(queue, task{node.Right, 2*index + 1})
			}
		}

		queue = queue[levelSize:]
	}

	return maxWidth
}
