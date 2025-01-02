struct Solution;

// https://leetcode.com/problems/product-of-array-except-self
impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut answer: Vec<i32> = vec![1; nums.len()];

        let mut prefix = 1;
        for i in 0..nums.len() {
            answer[i] = prefix;
            prefix *= nums[i];
        }

        let mut suffix = 1;
        for i in (0..nums.len()).rev() {
            answer[i] *= suffix;
            suffix *= nums[i];
        }

        answer
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::product_except_self(vec![1, 2, 3, 4]),
        vec![24, 12, 8, 6]
    );
    assert_eq!(
        Solution::product_except_self(vec![-1, 1, 0, -3, 3]),
        vec![0, 0, 9, 0, 0]
    );
}
