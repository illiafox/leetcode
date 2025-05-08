struct Solution;

impl Solution {
    // O(log n)
    pub fn find_peak_element(nums: Vec<i32>) -> i32 {
        let mut left = 0;
        let mut right = nums.len() - 1;

        while left < right {
            let mid = (left + right) / 2;
            if nums[mid] < nums[mid + 1] {
                left = mid + 1
            } else {
                right = mid
            }
        }

        left as i32
    }

    // O(n)
    pub fn find_peak_element_linear(nums: Vec<i32>) -> i32 {
        for i in 0..nums.len() {
            if (i == 0 || nums[i - 1] < nums[i]) && (i == nums.len() - 1 || nums[i + 1] < nums[i]) {
                return i as i32;
            }
        }

        0
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![1, 2, 3, 1], vec![2]),
        (vec![1, 2, 1, 3, 5, 6, 4], vec![5, 1]),
        (vec![1, 2], vec![1]),
    ];

    for (input, expected) in test_cases {
        assert!(
            expected.contains(&Solution::find_peak_element(input.clone())),
            "failed for input: {:?}",
            input,
        );
        assert!(
            expected.contains(&Solution::find_peak_element_linear(input.clone())),
            "failed for input: {:?}",
            input,
        );
    }
}
