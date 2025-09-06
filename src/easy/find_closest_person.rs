use std::cmp::Ordering::{Equal, Greater, Less};

struct Solution;

impl Solution {
    pub fn find_closest(x: i32, y: i32, z: i32) -> i32 {
        match (x - z).abs().cmp(&(y - z).abs()) {
            Equal => 0,
            Less => 1,
            Greater => 2,
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::easy::find_closest_person::Solution;

    #[test]
    fn test_find_closest() {
        let test_cases = [
            // (x, y, z, expected)
            (10, 15, 8, 1),    // x is closer
            (3, 8, 9, 2),      // y is closer
            (5, -5, 0, 0),     // tie
            (-7, -20, -10, 1), // x is closer with negatives
            (42, 100, 42, 1),  // x == z
            (100, 42, 42, 2),  // y == z
        ];

        for (x, y, z, expected) in test_cases {
            assert_eq!(
                Solution::find_closest(x, y, z),
                expected,
                "failed for x={x}, y={y}, z={z}"
            );
        }
    }
}
