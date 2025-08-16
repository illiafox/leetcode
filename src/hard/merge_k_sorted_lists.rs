use std::cmp::{Ordering, Reverse};
use std::collections::BinaryHeap;

struct Solution;

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

impl Ord for ListNode {
    fn cmp(&self, other: &Self) -> Ordering {
        self.val.cmp(&other.val)
    }
}
impl PartialOrd for ListNode {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Solution {
    // https://leetcode.com/problems/merge-k-sorted-lists
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        let mut dummy = Box::new(ListNode { val: 0, next: None });
        let mut cur = &mut dummy;

        let mut lists = lists;
        let mut heap = BinaryHeap::new();

        for item in &mut lists {
            if let Some(node) = item.take() {
                heap.push(Reverse(node));
            }
        }

        while let Some(Reverse(mut node)) = heap.pop() {
            let next = node.next.take();
            cur.next = Some(node);
            cur = cur.next.as_mut().unwrap();
            if let Some(n) = next {
                heap.push(Reverse(n));
            }
        }

        cur.next = None;
        dummy.next
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn list_from(v: &[i32]) -> Option<Box<ListNode>> {
        let mut head: Option<Box<ListNode>> = None;
        for &x in v.iter().rev() {
            head = Some(Box::new(ListNode { val: x, next: head }));
        }
        head
    }

    fn vec_from_list(mut head: Option<Box<ListNode>>) -> Vec<i32> {
        let mut out = Vec::new();
        while let Some(node) = head {
            out.push(node.val);
            head = node.next;
        }
        out
    }

    #[test]
    fn merge_k_lists_table() {
        let cases: Vec<(Vec<Vec<i32>>, Vec<i32>)> = vec![
            (vec![], vec![]),
            (vec![vec![], vec![], vec![]], vec![]),
            (vec![vec![1, 2, 3]], vec![1, 2, 3]),
            (
                vec![vec![1, 4, 5], vec![1, 3, 4], vec![2, 6]],
                vec![1, 1, 2, 3, 4, 4, 5, 6],
            ),
            (
                vec![vec![1, 1, 1], vec![1, 1], vec![1]],
                vec![1, 1, 1, 1, 1, 1],
            ),
            (
                vec![vec![-10, -1, 0], vec![-5, -2], vec![]],
                vec![-10, -5, -2, -1, 0],
            ),
            (
                vec![vec![5], vec![3], vec![4], vec![1], vec![2]],
                vec![1, 2, 3, 4, 5],
            ),
            (
                vec![
                    vec![],
                    vec![1],
                    vec![0, 2, 100],
                    vec![3, 4, 5, 6, 7],
                    vec![-1],
                ],
                vec![-1, 0, 1, 2, 3, 4, 5, 6, 7, 100],
            ),
        ];

        for (i, (lists_raw, expected)) in cases.into_iter().enumerate() {
            let lists: Vec<Option<Box<ListNode>>> =
                lists_raw.iter().map(|v| list_from(v)).collect();

            let got = vec_from_list(Solution::merge_k_lists(lists));
            assert_eq!(
                got, expected,
                "case {} failed: got {:?}, want {:?}",
                i, got, expected
            );
        }
    }
}
