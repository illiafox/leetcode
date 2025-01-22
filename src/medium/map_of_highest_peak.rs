use std::collections::VecDeque;

struct Solution;

impl Solution {
    pub fn highest_peak(is_water: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]; // Possible moves.

        let rows = is_water.len();
        let cols = is_water[0].len();

        let mut height = vec![vec![-1; cols]; rows];
        let mut queue = VecDeque::new();

        for r in 0..rows {
            for c in 0..cols {
                if is_water[r][c] == 1 {
                    height[r][c] = 0;
                    queue.push_back((r, c));
                }
            }
        }

        let mut next_layer_height = 1;

        while !queue.is_empty() {
            let current_queue_length = queue.len();

            for _ in 0..current_queue_length {
                let (i, j) = queue.pop_front().unwrap();

                for (dx, dy) in directions {
                    let new_i = i as i32 + dy;
                    let new_j = j as i32 + dx;

                    if new_j < 0 || new_j >= cols as i32 || new_i < 0 || new_i >= rows as i32 {
                        continue;
                    }

                    let new_i = new_i as usize;
                    let new_j = new_j as usize;

                    if height[new_i][new_j] != -1 {
                        continue;
                    }

                    height[new_i][new_j] = next_layer_height;
                    queue.push_back((new_i, new_j));
                }
            }

            next_layer_height += 1
        }

        height
    }
}
