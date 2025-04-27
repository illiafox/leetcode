package easy

func mergeTreesTraverse(parent **TreeNode, a, b *TreeNode) {
	if a == nil && b == nil {
		return
	}

	switch {
	case a == nil: // only b
		*parent = b
	case b == nil: // only a
		*parent = a
	default: // both present
		a.Val += b.Val // combine values
		*parent = a
	}

	left := func(n *TreeNode) *TreeNode {
		if n != nil {
			return n.Left
		}
		return nil
	}
	right := func(n *TreeNode) *TreeNode {
		if n != nil {
			return n.Right
		}
		return nil
	}

	mergeTreesTraverse(&(*parent).Left, left(a), left(b))
	mergeTreesTraverse(&(*parent).Right, right(a), right(b))
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var result *TreeNode
	mergeTreesTraverse(&result, root1, root2)

	return result
}
