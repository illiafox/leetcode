use std::cmp::{max, Ordering};

struct Solution;

impl Solution {
    pub fn longest_monotonic_subarray(nums: Vec<i32>) -> i32 {
        if nums.is_empty() {
            return 0;
        }

        let mut max_len = 1;
        let (mut incr_count, mut decr_count) = (1, 1);

        for window in nums.windows(2) {
            match window[1].cmp(&window[0]) {
                Ordering::Greater => {
                    incr_count += 1;
                    decr_count = 1;
                }
                Ordering::Less => {
                    decr_count += 1;
                    incr_count = 1;
                }
                Ordering::Equal => {
                    incr_count = 1;
                    decr_count = 1;
                }
            }
            max_len = max(max_len, incr_count.max(decr_count));
        }

        max_len
    }
}

fn test() {
    let test_cases = [
        (vec![1, 4, 3, 3, 2], 2),
        (vec![3, 3, 3, 3], 1),
        (vec![3, 2, 1], 3),
    ];

    for (nums, expected) in test_cases {
        assert_eq!(
            Solution::longest_monotonic_subarray(nums.clone()),
            expected,
            "failed for input: nums = {:?}",
            nums
        );
    }
}
