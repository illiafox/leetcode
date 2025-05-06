struct Solution;

impl Solution {
    pub fn check(nums: Vec<i32>) -> bool {
        if nums.len() <= 1 {
            return true;
        }

        let mut inversion_count = 0;

        for i in 1..nums.len() {
            if nums[i] < nums[i - 1] {
                inversion_count += 1;
            }
        }

        if nums[0] < nums[nums.len() - 1] {
            inversion_count += 1;
        }

        inversion_count <= 1
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![3, 4, 5, 1, 2], true),
        (vec![2, 1, 3, 4], false),
        (vec![1, 2, 3], true),
    ];

    for (nums, expected) in test_cases {
        assert_eq!(
            Solution::check(nums.clone()),
            expected,
            "failed for input: nums = {:?}",
            nums
        );
    }
}
