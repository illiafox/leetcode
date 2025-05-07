use std::cmp::{max, min, Ordering};

struct Solution;

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut start = 0;
        let mut end = height.len() - 1;

        let mut max_area = 0;

        while start < end {
            let area = min(height[start], height[end]) * (end - start) as i32;
            max_area = max(max_area, area);

            match height[start].cmp(&height[end]) {
                Ordering::Less | Ordering::Equal => start += 1,
                Ordering::Greater => end -= 1,
            }
        }

        max_area
    }
}

#[test]
fn test() {
    let test_cases = [(vec![1, 8, 6, 2, 5, 4, 8, 3, 7], 49), (vec![1, 1], 1)];

    for (height, expected) in test_cases {
        assert_eq!(
            Solution::max_area(height.clone()),
            expected,
            "failed for input: {:?}",
            height
        );
    }
}
