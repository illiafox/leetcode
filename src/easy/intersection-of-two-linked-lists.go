package easy

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visited := map[*ListNode]struct{}{}

	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := visited[headA]; ok {
				return headA
			}
			visited[headA] = struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := visited[headB]; ok {
				return headB
			}
			visited[headB] = struct{}{}
			headB = headB.Next
		}
		if headA == headB {
			return headA
		}
	}

	return nil
}
