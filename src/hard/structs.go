package hard

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}

	i := 1
	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// left child
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Left)
			}
			i++
		}

		// right child
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func listFrom(vals []int) *ListNode {
	var head, tail *ListNode
	for _, v := range vals {
		n := &ListNode{Val: v}
		if head == nil {
			head, tail = n, n
		} else {
			tail.Next = n
			tail = n
		}
	}
	return head
}

func sliceFromList(head *ListNode) []int {
	var out []int
	for p := head; p != nil; p = p.Next {
		out = append(out, p.Val)
	}
	return out
}
