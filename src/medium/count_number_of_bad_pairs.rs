use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn count_bad_pairs(nums: Vec<i32>) -> i64 {
        let n = nums.len() as i64;
        let total_pairs = n * (n - 1) / 2; // sum
        let mut good_pairs = 0;

        let mut m: HashMap<i32, i32> = HashMap::new();

        for (i, &num) in nums.iter().enumerate() {
            *m.entry(num - i as i32).or_default() += 1
        }

        for (i, &num) in nums.iter().enumerate() {
            if let Some(m) = m.get_mut(&(num - i as i32)) {
                *m -= 1;
                if *m > 0 {
                    good_pairs += *m as i64;
                }
            }
        }

        total_pairs - good_pairs
    }
}

#[test]
fn test() {
    let test_cases = [(vec![4, 1, 3, 3], 5), (vec![1, 2, 3, 4, 5], 0)];

    for (nums, expected) in test_cases {
        assert_eq!(
            Solution::count_bad_pairs(nums.clone()),
            expected,
            "failed for nums: {:?}",
            nums
        );
    }
}
