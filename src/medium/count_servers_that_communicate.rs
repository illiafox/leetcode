struct Solution;

impl Solution {
    pub fn count_servers(grid: Vec<Vec<i32>>) -> i32 {
        let rows = grid.len();
        let cols = grid[0].len();

        let mut rows_computers = vec![0; rows];
        let mut cols_computers = vec![0; cols];

        let mut total_computers = 0;

        for i in 0..rows {
            for (j, &is_computer) in grid[i].iter().enumerate() {
                if is_computer == 0 {
                    continue;
                }

                rows_computers[i] += 1;
                cols_computers[j] += 1;
                total_computers += 1;
            }
        }

        for i in 0..rows {
            for (j, &is_computer) in grid[i].iter().enumerate() {
                if is_computer == 0 {
                    continue;
                }

                if rows_computers[i] == 1 && cols_computers[j] == 1 {
                    total_computers -= 1
                }
            }
        }

        total_computers
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![vec![1, 0], vec![0, 1]], 0),
        (vec![vec![1, 0], vec![1, 1]], 3),
        (
            vec![
                vec![1, 1, 0, 0],
                vec![0, 0, 1, 0],
                vec![0, 0, 1, 0],
                vec![0, 0, 0, 1],
            ],
            4,
        ),
        (
            vec![
                vec![1, 0, 0, 1, 0],
                vec![0, 0, 0, 0, 0],
                vec![0, 0, 0, 1, 0],
            ],
            3,
        ),
    ];

    for (grid, expected) in test_cases {
        assert_eq!(
            Solution::count_servers(grid.clone()),
            expected,
            "failed for input grid: {:?}",
            grid
        );
    }
}
