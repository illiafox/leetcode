use std::collections::{HashMap, VecDeque};

struct Solution;

#[derive(PartialEq, Eq, Clone, Debug)]
struct Task {
    coord: (usize, usize),
    energy: i32,
    collected_litter: i32,
    visited_collectibles: i16,
}

// https://leetcode.com/problems/minimum-moves-to-clean-the-classroom
impl Solution {
    pub fn min_moves(classroom: Vec<String>, energy: i32) -> i32 {
        const STUDENT_CELL: char = 'S';
        const LITTER_CELL: char = 'L';
        const RESET_CELL: char = 'R';
        const OBSTACLE_CELL: char = 'X';
        const EMPTY_CELL: char = '.';

        let classroom_chars: Vec<Vec<char>> =
            classroom.iter().map(|s| s.chars().collect()).collect();

        let mut total_litter = 0;
        let mut litter_by_id = HashMap::new();

        let mut start: Option<(usize, usize)> = None;

        for (i, row) in classroom_chars.iter().enumerate() {
            for (j, &c) in row.iter().enumerate() {
                if c == STUDENT_CELL {
                    start = Some((i, j))
                }
                if c == LITTER_CELL {
                    total_litter += 1;
                    litter_by_id.insert((i, j), total_litter);
                }
            }
        }
        // https://www.geeksforgeeks.org/dsa/what-is-bitmasking/
        if start.is_none() {
            panic!("start not found")
        }

        if total_litter == 0 {
            return 0;
        }

        let coord = start.unwrap();

        let mut queue = VecDeque::from([Task {
            coord,
            energy,
            visited_collectibles: 0,
            collected_litter: 0,
        }]);

        let mut best_effort: HashMap<(usize, usize, i16), i32> = HashMap::new();

        let directions = [(1, 0), (-1, 0), (0, 1), (0, -1)];
        let (rows, cols) = (classroom.len(), classroom[0].len());

        let mut step = 0;

        while !queue.is_empty() {
            let size = queue.len();
            for _ in 0..size {
                let mut task = queue.pop_front().unwrap();
                let (i, j) = task.coord;

                let c = classroom_chars[i][j];
                if c == RESET_CELL {
                    task.energy = energy
                }

                if c == LITTER_CELL {
                    let lit_id = litter_by_id.get(&(i, j)).unwrap();

                    if task.visited_collectibles & (1 << lit_id) == 0 {
                        task.visited_collectibles |= 1 << lit_id;

                        task.collected_litter += 1;
                        if task.collected_litter == total_litter {
                            return step;
                        }
                    }
                }

                if task.energy <= 0 {
                    continue;
                }

                for (dx, dy) in directions {
                    let new_i = i as i32 + dy;
                    let new_j = j as i32 + dx;

                    if new_j < 0 || new_j >= cols as i32 || new_i < 0 || new_i >= rows as i32 {
                        continue;
                    }

                    let new_i = new_i as usize;
                    let new_j = new_j as usize;

                    if classroom_chars[new_i][new_j] != OBSTACLE_CELL {
                        let key = (new_i, new_j, task.visited_collectibles);
                        if let Some(&max_energy) = best_effort.get(&key)
                            && max_energy >= task.energy - 1
                        {
                            continue;
                        }
                        best_effort.insert(key, task.energy - 1);

                        let mut t_clone = task.clone();
                        t_clone.coord = (new_i, new_j);
                        t_clone.energy = task.energy - 1;

                        queue.push_back(t_clone)
                    }
                }
            }

            step += 1;
        }

        -1
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_min_moves() {
        let cases: Vec<(&[&str], i32, i32)> = vec![
            (&["S.", "XL"], 2, 2),
            (&["LS", "RL"], 4, 3),
            (&["L.S", "RXL"], 3, -1),
            (&["RRLXX", "X.RSL"], 8, 4),
            (&["S.", ".."], 5, 0),
            (&["S.X", "XX.", "X.L"], 10, -1),
            (&["S...", "...L"], 3, -1),
            (&["S.R.L"], 2, 4),
        ];

        for (grid, energy, expected) in cases {
            let classroom = grid.iter().map(|&s| s.to_string()).collect();
            assert_eq!(Solution::min_moves(classroom, energy), expected);
        }
    }
}
