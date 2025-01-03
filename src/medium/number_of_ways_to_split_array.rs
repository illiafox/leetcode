struct Solution;

// https://leetcode.com/problems/number-of-ways-to-split-array
impl Solution {
    pub fn ways_to_split_array(nums: Vec<i32>) -> i32 {
        let mut right_sum: i64 = nums.iter().map(|&x| x as i64).sum();
        let mut left_sum: i64 = 0;

        let mut valid_splits = 0;

        for &n in &nums[..nums.len() - 1] {
            left_sum += n as i64;
            right_sum -= n as i64;

            if left_sum >= right_sum {
                valid_splits += 1;
            }
        }

        valid_splits
    }
}

#[test]
fn test() {
    assert_eq!(Solution::ways_to_split_array(vec![10, 4, -8, 7]), 2);
    assert_eq!(Solution::ways_to_split_array(vec![2, 3, 1, 0]), 2);
}
