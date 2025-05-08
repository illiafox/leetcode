use std::cmp::Reverse;
use std::collections::BinaryHeap;

struct Solution;

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut heap = BinaryHeap::new();

        for n in nums {
            heap.push(Reverse(n));
            if heap.len() > k as usize {
                heap.pop();
            }
        }

        heap.peek().unwrap().0
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![3, 2, 1, 5, 6, 4], 2, 5),
        (vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4, 4),
    ];

    for (nums, k, expected) in test_cases {
        assert_eq!(
            Solution::find_kth_largest(nums.clone(), k),
            expected,
            "failed for nums {:?} k {}",
            nums,
            k,
        );
    }
}
