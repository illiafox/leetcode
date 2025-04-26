package easy

// returns depth and parent of the cousin
func isCousinsTraverse(root, prev *TreeNode, depth, target int) (*int, *TreeNode) {
	if root == nil {
		return nil, nil
	}

	if root.Val == target {
		return &depth, prev
	}

	depth++
	for _, node := range []*TreeNode{root.Right, root.Left} {
		if res, resNode := isCousinsTraverse(node, root, depth, target); res != nil {
			return res, resNode
		}
	}

	return nil, nil
}

func isCousins(root *TreeNode, x int, y int) bool {
	xDepth, xNode := isCousinsTraverse(root, nil, 0, x)
	yDepth, yNode := isCousinsTraverse(root, nil, 0, y)

	return xDepth != nil && yDepth != nil && *xDepth == *yDepth && xNode != yNode
}
