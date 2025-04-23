use std::collections::{HashSet, VecDeque};

struct Solution;

impl Solution {
    pub fn oranges_rotting(grid: Vec<Vec<i32>>) -> i32 {
        const FRESH: i32 = 1;
        const ROTTEN: i32 = 2;

        let directions = [(1, 0), (-1, 0), (0, 1), (0, -1)];
        let (rows, cols) = (grid.len(), grid[0].len());

        let mut fresh = HashSet::new();
        let mut queue = VecDeque::new();

        for (i, row) in grid.iter().enumerate() {
            for j in 0..row.len() {
                match grid[i][j] {
                    FRESH => {
                        fresh.insert((i, j));
                    }
                    ROTTEN => {
                        queue.push_back((i, j));
                    }
                    _ => {}
                }
            }
        }

        if fresh.is_empty() {
            return 0;
        }
        if queue.is_empty() {
            return -1;
        }

        let mut minutes = 0;

        while !fresh.is_empty() {
            let mut changed = false;
            let level_size = queue.len();

            for _ in 0..level_size {
                let (i, j) = queue.pop_front().unwrap();

                for (di, dj) in directions {
                    let ni = i as isize + di;
                    let nj = j as isize + dj;

                    if ni >= 0 && ni < rows as isize && nj >= 0 && nj < cols as isize {
                        let (ni, nj) = (ni as usize, nj as usize);
                        if fresh.remove(&(ni, nj)) {
                            queue.push_back((ni, nj));
                            changed = true;
                        }
                    }
                }
            }

            if !changed {
                return -1;
            }

            minutes += 1;
        }

        minutes
    }
}
