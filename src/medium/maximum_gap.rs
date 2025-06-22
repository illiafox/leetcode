struct Solution;

#[derive(Clone)]
struct Bucket {
    min: Option<i32>,
    max: Option<i32>,
}

// https://leetcode.com/problems/maximum-gap
impl Solution {
    pub fn maximum_gap(nums: Vec<i32>) -> i32 {
        let n = nums.len() as i32;

        if n < 2 {
            return 0;
        }

        let max_val = nums.iter().max().unwrap();
        let min_val = nums.iter().min().unwrap();

        if min_val == max_val {
            return 0;
        }

        // (a + b - 1) / b is integer math for ceil(a / b)
        let bucket_size = ((max_val - min_val) + (n - 2)) / (n - 1);
        let bucket_count = (max_val - min_val) / bucket_size + 1;

        let mut buckets = vec![
            Bucket {
                min: None,
                max: None
            };
            bucket_count as usize
        ];

        for &number in &nums {
            let index = ((number - min_val) / bucket_size) as usize;
            let bucket = &mut buckets[index];
            bucket.min = Some(bucket.min.map_or(number, |min| min.min(number)));
            bucket.max = Some(bucket.max.map_or(number, |max| max.max(number)));
        }

        let mut max_diff = 0;

        let mut prev_max = None;
        for bucket in buckets {
            if let (Some(min), Some(max)) = (bucket.min, bucket.max) {
                if let Some(prev) = prev_max {
                    max_diff = max_diff.max(min - prev);
                }
                prev_max = Some(max);
            }
        }

        max_diff
    }

    pub fn maximum_gap_inefficient(nums: Vec<i32>) -> i32 {
        if nums.len() < 2 {
            return 0;
        }

        let mut nums = nums;
        nums.sort();

        let mut min_diff = 0;

        for window in nums.windows(2) {
            min_diff = min_diff.max(window[1] - window[0]);
        }

        min_diff
    }
}

#[test]
fn test() {
    let test_cases = [(vec![3, 6, 9, 1], 3), (vec![10], 0), (vec![0, 0, 0], 0)];

    for (nums, expected) in test_cases {
        assert_eq!(
            Solution::maximum_gap(nums.clone()),
            expected,
            "failed for nums {nums:?}",
        );
        assert_eq!(
            Solution::maximum_gap_inefficient(nums.clone()),
            expected,
            "failed for nums {nums:?}",
        );
    }
}
