package easy

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode

	for start := head; start != nil; {
		if start.Val == val {
			if prev == nil {
				start = head.Next
				head = head.Next
				prev = nil
				continue
			}

			prev.Next = start.Next
		} else {
			prev = start
		}

		start = start.Next
	}

	return head
}
