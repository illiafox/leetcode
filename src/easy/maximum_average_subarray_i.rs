struct Solution;

impl Solution {
    pub fn find_max_average(nums: Vec<i32>, k: i32) -> f64 {
        let mut s: i32 = nums.iter().take(k as usize).sum();

        let mut max_avg = s as f64 / k as f64;

        for i in k as usize..nums.len() {
            s -= nums[i - k as usize];
            s += nums[i];

            let avg = s as f64 / k as f64;
            if max_avg < avg {
                max_avg = avg
            }
        }

        max_avg
    }

    pub fn find_max_average_ineffective(nums: Vec<i32>, k: i32) -> f64 {
        let mut max_avg = f64::MIN;

        for window in nums.windows(k as usize) {
            let avg = window.iter().sum::<i32>() as f64 / k as f64;

            if max_avg < avg {
                max_avg = avg
            }
        }

        max_avg
    }
}

#[test]
fn test() {
    let test_cases = [(vec![1, 12, -5, -6, 50, 3], 4, 12.75000), (vec![5], 1, 5.0)];

    for (nums, k, expected) in test_cases {
        assert_eq!(
            Solution::find_max_average(nums.clone(), k),
            expected,
            "failed for input: nums = {:?}",
            nums
        );
        assert_eq!(
            Solution::find_max_average_ineffective(nums.clone(), k),
            expected,
            "failed for input: nums = {:?}",
            nums
        );
    }
}
