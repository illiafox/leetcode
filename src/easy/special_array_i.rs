struct Solution;

impl Solution {
    pub fn is_array_special(nums: Vec<i32>) -> bool {
        for (i, num) in nums.iter().skip(1).enumerate() {
            if nums[i] & 1 == num & 1 {
                return false;
            }
        }

        true
    }

    pub fn is_array_special_windows(nums: Vec<i32>) -> bool {
        nums.windows(2).all(|pair| (pair[0] & 1) != (pair[1] & 1))
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![1, 2], true),
        (vec![2, 1, 4], true),
        (vec![4, 3, 1, 6], false),
    ];

    for (nums, expected) in test_cases {
        assert_eq!(
            Solution::is_array_special(nums.clone()),
            expected,
            "failed for input: nums = {:?}",
            nums
        );
        assert_eq!(
            Solution::is_array_special_windows(nums.clone()),
            expected,
            "failed for input: nums = {:?}",
            nums
        );
    }
}
