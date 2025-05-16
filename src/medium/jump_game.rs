struct Solution;

impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut farthest = 0;

        for (i, &num) in nums.iter().enumerate() {
            if i > farthest {
                return false;
            }
            farthest = farthest.max(i + num as usize);
        }

        true
    }

    pub fn can_jump_dp(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut dp = vec![false; n];
        dp[0] = true;

        for i in 1..n {
            for j in 0..i {
                if dp[j] && j + nums[j] as usize >= i {
                    dp[i] = true;
                    break;
                }
            }
        }

        dp[n - 1]
    }
}

#[test]
fn test() {
    // https://leetcode.com/problems/jump-game

    let test_cases = [(vec![2, 3, 1, 1, 4], true), (vec![3, 2, 1, 0, 4], false)];

    for (nums, expected) in test_cases {
        assert_eq!(
            Solution::can_jump(nums.clone()),
            expected,
            "failed for nums: {nums:?}",
        );
        assert_eq!(
            Solution::can_jump_dp(nums.clone()),
            expected,
            "failed for nums: {nums:?}",
        );
    }
}
