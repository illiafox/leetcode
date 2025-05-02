package medium

func reverseOddLevels(root *TreeNode) *TreeNode {
	level := 0

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)

		if level&1 == 1 {
			for i, j := 0, levelSize-1; i < j; i, j = i+1, j-1 {
				queue[i].Val, queue[j].Val = queue[j].Val, queue[i].Val
			}
		}

		if queue[0].Right == nil && queue[0].Left == nil {
			return root
		}

		for i := 0; i < levelSize; i++ {
			queue = append(queue, queue[i].Right, queue[i].Left)
		}

		queue = queue[levelSize:]
		level++
	}

	return root
}
