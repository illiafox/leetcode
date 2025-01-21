use std::cmp::{max, min};

struct Solution;

impl Solution {
    pub fn grid_game(grid: Vec<Vec<i32>>) -> i64 {
        let mut up_sum: i64 = grid[0].iter().map(|&x| x as i64).sum();

        let mut left_sum = 0;

        let mut min_val: i64 = up_sum;

        for (i, &num) in grid[0].iter().enumerate() {
            up_sum -= num as i64;

            min_val = min(min_val, max(up_sum, left_sum));
            left_sum += grid[1][i] as i64;
        }

        min_val
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![vec![2, 5, 4], vec![1, 5, 1]], 4),
        (vec![vec![3, 3, 1], vec![8, 5, 2]], 4),
        (vec![vec![1, 3, 1, 15], vec![1, 3, 3, 1]], 7),
    ];

    for (input, expected) in test_cases {
        assert_eq!(
            Solution::grid_game(input.clone()),
            expected,
            "failed for input: {:?}",
            input,
        );
    }
}
