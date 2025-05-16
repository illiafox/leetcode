struct Solution;

impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        let rows = obstacle_grid.len();
        let cols = obstacle_grid[0].len();

        let mut dp = vec![vec![0; cols]; rows];
        dp[0][0] = 1;

        for i in 0..rows {
            for j in 0..cols {
                if obstacle_grid[i][j] == 1 {
                    dp[i][j] = 0;
                    continue;
                }

                if j > 0 {
                    dp[i][j] += dp[i][j - 1]
                }
                if i > 0 {
                    dp[i][j] += dp[i - 1][j]
                }
            }
        }

        dp[rows - 1][cols - 1]
    }
}

#[test]
fn test() {
    // https://leetcode.com/problems/unique-paths-ii/

    let test_cases = [
        (vec![vec![0, 0, 0], vec![0, 1, 0], vec![0, 0, 0]], 2),
        (vec![vec![0, 1], vec![0, 0]], 1),
    ];

    for (matrix, expected) in test_cases {
        assert_eq!(
            Solution::unique_paths_with_obstacles(matrix.clone()),
            expected,
            "failed for matrix: {matrix:?}",
        );
    }
}
