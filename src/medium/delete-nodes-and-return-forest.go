package medium

func delNodesTraverse(root *TreeNode, toDelete map[int]struct{}, res *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = delNodesTraverse(root.Left, toDelete, res)
	root.Right = delNodesTraverse(root.Right, toDelete, res)

	if _, ok := toDelete[root.Val]; ok {
		if root.Left != nil {
			*res = append(*res, root.Left)
		}
		if root.Right != nil {
			*res = append(*res, root.Right)
		}
		return nil
	}

	return root
}

// https://leetcode.com/problems/delete-nodes-and-return-forest/
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	m := make(map[int]struct{}, len(toDelete))
	for _, v := range toDelete {
		m[v] = struct{}{}
	}

	res := make([]*TreeNode, 0, 1)
	if node := delNodesTraverse(root, m, &res); node != nil {
		res = append(res, node)
	}
	return res
}
