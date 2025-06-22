struct Solution;

use std::cmp::Reverse;
use std::collections::{BinaryHeap, HashSet};

impl Solution {
    pub fn third_max(nums: Vec<i32>) -> i32 {
        let mut heap = BinaryHeap::new();
        let mut seen = HashSet::new();

        for &n in &nums {
            if seen.insert(n) {
                heap.push(Reverse(n));
                if heap.len() > 3 {
                    heap.pop();
                }
            }
        }

        if heap.len() == 3 {
            return heap.peek().unwrap().0;
        }

        heap.iter().map(|r| r.0).max().unwrap()
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![3, 2, 1], 1),
        (vec![2, 2, 3, 1], 1),
        (vec![1, 2, 2, 5, 3, 5], 2),
    ];

    for (nums, expected) in test_cases {
        assert_eq!(
            Solution::third_max(nums.clone()),
            expected,
            "failed for nums {nums:?}"
        );
    }
}
