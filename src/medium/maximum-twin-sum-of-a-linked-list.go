package medium

func pairSum(head *ListNode) int {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reverseList := func(start *ListNode) *ListNode {
		cur := start
		var prev *ListNode
		var next *ListNode
		for cur != nil {
			next = cur.Next
			cur.Next = prev
			prev, cur = cur, next
		}
		return prev
	}
	slow = reverseList(slow)

	maxSum := 0
	for slow != nil && head != nil {
		maxSum = max(maxSum, slow.Val+head.Val)
		slow, head = slow.Next, head.Next
	}

	return maxSum
}
