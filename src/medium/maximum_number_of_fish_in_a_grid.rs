use std::cmp::max;
struct Solution;

impl Solution {
    pub fn dfs(grid: &Vec<Vec<i32>>, (i, j): (usize, usize), visited: &mut Vec<Vec<bool>>) -> i32 {
        let directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]; // Possible moves.

        if visited[i][j] {
            return 0;
        }
        visited[i][j] = true;

        let mut sum = grid[i][j];

        for (dx, dy) in directions {
            let new_i = i as i32 + dy;
            let new_j = j as i32 + dx;

            if new_j < 0 || new_j >= grid[0].len() as i32 || new_i < 0 || new_i >= grid.len() as i32
            {
                continue;
            }

            let new_i = new_i as usize;
            let new_j = new_j as usize;

            if grid[new_i][new_j] == 0 {
                continue;
            }

            sum += Self::dfs(grid, (new_i, new_j), visited)
        }

        sum
    }

    pub fn find_max_fish(grid: Vec<Vec<i32>>) -> i32 {
        let mut max_fish = 0;

        let mut visited = vec![vec![false; grid[0].len()]; grid.len()];

        for i in 0..grid.len() {
            for j in 0..grid[i].len() {
                if grid[i][j] > 0 {
                    max_fish = max(max_fish, Self::dfs(&grid, (i, j), &mut visited))
                }
            }
        }

        max_fish
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            vec![
                vec![0, 2, 1, 0],
                vec![4, 0, 0, 3],
                vec![1, 0, 0, 4],
                vec![0, 3, 2, 0],
            ],
            7,
        ),
        (
            vec![
                vec![1, 0, 0, 0],
                vec![0, 0, 0, 0],
                vec![0, 0, 0, 0],
                vec![0, 0, 0, 1],
            ],
            1,
        ),
    ];

    for (grid, expected) in test_cases {
        assert_eq!(
            Solution::find_max_fish(grid.clone()),
            expected,
            "failed for grid: {:?}",
            grid,
        );
    }
}
