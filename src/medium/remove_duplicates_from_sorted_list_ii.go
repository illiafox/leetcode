package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	out := &ListNode{Next: head}

	prev := out

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}

			prev.Next = head.Next
		} else {
			prev = prev.Next
		}

		head = head.Next
	}

	return out.Next
}

func deleteDuplicatesOld(head *ListNode) *ListNode {
	var beforePrev *ListNode
	var prev *ListNode

	current := head

	removed := make(map[int]struct{})

	for current != nil {
		if prev != nil && current.Val == prev.Val {
			removed[current.Val] = struct{}{}

			if beforePrev != nil {
				beforePrev.Next = current.Next

				current = beforePrev

				prev = nil
				beforePrev = nil
			} else {

				current = current.Next
				head = current

				prev = nil
				beforePrev = nil
			}
		} else if _, ok := removed[current.Val]; ok {
			if prev != nil {
				prev.Next = current.Next
				current = current.Next

				beforePrev = nil
			} else {
				head = current.Next
				current = current.Next

				prev = nil
				beforePrev = nil
			}
		} else {
			beforePrev = prev
			prev = current
			current = current.Next
		}
	}

	return head
}
