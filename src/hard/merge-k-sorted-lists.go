package hard

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.com/problems/merge-k-sorted-lists/
// better solution with minheap is written in rust (see merge_k_sorted_lists.rs)
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for {
		minVal := math.MaxInt

		atLeastOne := false
		for _, list := range lists {
			if list == nil {
				continue
			}
			atLeastOne = true

			minVal = min(minVal, list.Val)
		}

		if !atLeastOne {
			break
		}

		for i, list := range lists {
			if list == nil {
				continue
			}

			if list.Val == minVal {
				node := &ListNode{
					Val:  list.Val,
					Next: nil,
				}
				cur.Next = node
				cur = node
				lists[i] = lists[i].Next
			}
		}

	}

	return dummy.Next
}
