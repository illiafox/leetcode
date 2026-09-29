use std::collections::VecDeque;

struct Solution;

#[derive(Clone)]
struct Task {
    x: usize,
    y: usize,
    brackets: i32,
}

impl Solution {
    pub fn has_valid_path_dynamic(grid: Vec<Vec<char>>) -> bool {
        let (rows, cols) = (grid.len(), grid[0].len());
        if grid[0][0] == ')' || grid[rows - 1][cols - 1] == '(' || (rows + cols - 1) % 2 != 0 {
            return false;
        }

        let mut dp = vec![0u128; cols];

        for i in 0..rows {
            for j in 0..cols {
                let prev_mask = if i == 0 && j == 0 {
                    1u128
                } else {
                    let from_top = if i > 0 { dp[j] } else { 0 };
                    let from_left = if j > 0 { dp[j - 1] } else { 0 };
                    from_top | from_left
                };

                dp[j] = match grid[i][j] {
                    '(' => prev_mask << 1,
                    ')' => prev_mask >> 1,
                    _ => 0,
                };
            }
        }

        dp[cols - 1] & 1 != 0
    }

    pub fn has_valid_path(grid: Vec<Vec<char>>) -> bool {
        let check = |x| match x {
            '(' => 1,
            ')' => -1,
            _ => panic!("unknown char {}", x),
        };

        let (rows, cols) = (grid.len(), grid[0].len());
        if grid[0][0] == ')' || grid[rows - 1][cols - 1] == '(' || (rows + cols - 1) % 2 != 0 {
            return false;
        }

        let mut visited = vec![vec![0u128; cols]; rows];
        let max_balance = ((rows + cols) / 2) as i32;

        let mut queue = VecDeque::new();
        queue.push_front(Task {
            x: 0,
            y: 0,
            brackets: 0,
        });

        while !queue.is_empty() {
            let l = queue.len();

            for _ in 0..l {
                let mut t = queue.pop_front().unwrap();

                t.brackets += check(grid[t.y][t.x]);
                if t.brackets < 0 || t.brackets > max_balance {
                    continue;
                }

                let shift = 1u128 << t.brackets;
                if visited[t.y][t.x] & shift > 0 {
                    continue;
                }
                visited[t.y][t.x] |= shift;

                if t.y == rows - 1 && t.x == cols - 1 {
                    if t.brackets == 0 {
                        return true;
                    }

                    continue;
                }

                for (di, dj) in [(1, 0), (0, 1)] {
                    let ny = t.y + di;
                    let nx = t.x + dj;

                    if ny < rows && nx < cols {
                        queue.push_back(Task {
                            x: nx,
                            y: ny,
                            brackets: t.brackets,
                        });
                    }
                }
            }
        }

        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn make_grid(slices: &[&str]) -> Vec<Vec<char>> {
        slices.iter().map(|row| row.chars().collect()).collect()
    }

    fn check_both(grid: &[&str], expected: bool) {
        let g1 = make_grid(grid);
        let g2 = g1.clone();

        assert_eq!(
            Solution::has_valid_path(g1),
            expected,
            "Failed on has_valid_path"
        );
        assert_eq!(
            Solution::has_valid_path_dynamic(g2),
            expected,
            "Failed on has_valid_path_dynamic"
        );
    }

    #[test]
    fn test_leetcode_description_example_1() {
        let grid = ["(((", ")()", "()(", "())"];
        check_both(&grid, true);
    }

    #[test]
    fn test_leetcode_description_example_2() {
        let grid = ["))", "(("];
        check_both(&grid, false);
    }

    #[test]
    fn test_large_valid_100x99_grid() {
        let rows = 100;
        let cols = 99;
        let mut grid_strings = Vec::with_capacity(rows);

        for r in 0..rows {
            let mut row = String::with_capacity(cols);
            for c in 0..cols {
                if r == 0 {
                    row.push('(');
                } else if c == cols - 1 {
                    row.push(')');
                } else {
                    // Checkerboard pattern creates maximum path branching
                    row.push(if (r + c) % 2 == 0 { '(' } else { ')' });
                }
            }
            grid_strings.push(row);
        }

        let slice: Vec<&str> = grid_strings.iter().map(|s| s.as_str()).collect();
        check_both(&slice, true);
    }

    #[test]
    fn test_large_all_open_top_all_close_bottom_100x99() {
        let rows = 100;
        let cols = 99;
        let mut grid_strings = Vec::with_capacity(rows);

        for r in 0..rows {
            let ch = if r < rows / 2 { '(' } else { ')' };
            grid_strings.push(ch.to_string().repeat(cols));
        }

        let slice: Vec<&str> = grid_strings.iter().map(|s| s.as_str()).collect();
        check_both(&slice, true);
    }

    #[test]
    fn test_large_100x100_odd_length_fail() {
        let rows = 100;
        let cols = 100;
        let grid_strings: Vec<String> = (0..rows)
            .map(|r| {
                (0..cols)
                    .map(|c| if (r + c) % 2 == 0 { '(' } else { ')' })
                    .collect()
            })
            .collect();

        let slice: Vec<&str> = grid_strings.iter().map(|s| s.as_str()).collect();
        check_both(&slice, false);
    }

    #[test]
    fn test_large_worst_case_pruned_grid() {
        let rows = 80;
        let cols = 81;
        let mut grid_strings = Vec::with_capacity(rows);

        for r in 0..rows {
            let mut row = String::with_capacity(cols);
            for c in 0..cols {
                if r == rows - 1 && c == cols - 1 {
                    row.push(')');
                } else if r == 0 && c == 0 {
                    row.push('(');
                } else if r + c > 120 {
                    row.push(')');
                } else {
                    row.push('(');
                }
            }
            grid_strings.push(row);
        }

        let slice: Vec<&str> = grid_strings.iter().map(|s| s.as_str()).collect();
        check_both(&slice, false);
    }

    #[test]
    fn test_static_slice_20x21() {
        let grid = [
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            "(((((((((((((((((((((",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
            ")))))))))))))))))))))",
        ];
        check_both(&grid, true);
    }
}
