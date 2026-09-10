package medium

func swapPairsSimplified(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		first.Next = second.Next
		second.Next = first
		prev.Next = second

		prev = first
	}

	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	var next *ListNode
	swapNodes(head, &next)

	return next
}

func swapNodes(head *ListNode, set **ListNode) {
	prev, head := head, head.Next

	for pair := true; head != nil; pair = !pair {
		if pair {
			next := head.Next
			head.Next = prev
			prev.Next = next
			*set = head

			head = next
		} else {
			set = &prev.Next
			prev = head
			head = head.Next
		}
	}
}
