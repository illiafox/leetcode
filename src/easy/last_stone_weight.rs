struct Solution;

use std::collections::BinaryHeap;

impl Solution {
    pub fn last_stone_weight(stones: Vec<i32>) -> i32 {
        let mut heap = BinaryHeap::from(stones);

        while heap.len() > 1 {
            if let (Some(a), Some(b)) = (heap.pop(), heap.pop()) {
                if a == b {
                    continue;
                }

                heap.push(a - b)
            }
        }

        *heap.peek().unwrap_or(&0)
    }
}

#[test]
fn test() {
    let test_cases = [(vec![2, 7, 4, 1, 8, 1], 1), (vec![1], 1), (vec![2, 2], 0)];

    for (input, expected) in test_cases {
        assert_eq!(
            Solution::last_stone_weight(input.clone()),
            expected,
            "failed for input: nums = {:?}",
            input
        );
    }
}
