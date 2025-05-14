use std::cmp::min;

struct Solution;

impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let rows = grid.len();
        let cols = grid[0].len();

        let mut dp = vec![vec![0; cols]; rows];

        dp[0][0] = grid[0][0];
        for j in 1..cols {
            dp[0][j] = grid[0][j] + dp[0][j - 1]
        }
        for i in 1..rows {
            dp[i][0] = grid[i][0] + dp[i - 1][0]
        }

        for i in 1..rows {
            for j in 1..cols {
                dp[i][j] = min(dp[i][j - 1], dp[i - 1][j]) + grid[i][j]
            }
        }

        dp[rows - 1][cols - 1]
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1]], 7),
        (vec![vec![1, 2, 3], vec![4, 5, 6]], 12),
    ];

    for (grid, expected) in test_cases {
        assert_eq!(
            Solution::min_path_sum(grid.clone()),
            expected,
            "failed for grid: {:?}",
            grid,
        );
    }
}
