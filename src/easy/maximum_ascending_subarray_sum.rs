use std::cmp::{max, Ordering};

struct Solution;

impl Solution {
    pub fn max_ascending_sum(nums: Vec<i32>) -> i32 {
        let mut max_sum = nums[0];
        let mut curr_sum = max_sum;

        for window in nums.windows(2) {
            match window[1].cmp(&window[0]) {
                Ordering::Greater => curr_sum += window[1],
                Ordering::Less | Ordering::Equal => {
                    max_sum = max_sum.max(curr_sum);
                    curr_sum = window[1]
                }
            }
        }

        max(max_sum, curr_sum)
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![10, 20, 30, 5, 10, 50], 65),
        (vec![10, 20, 30, 40, 50], 150),
        (vec![12, 17, 15, 13, 10, 11, 12], 33),
        (vec![3, 6, 10, 1, 8, 9, 9, 8, 9], 19),
    ];

    for (nums, expected) in test_cases {
        assert_eq!(
            Solution::max_ascending_sum(nums.clone()),
            expected,
            "failed for input: nums = {:?}",
            nums
        );
    }
}
