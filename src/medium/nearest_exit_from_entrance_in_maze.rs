use std::collections::{HashSet, VecDeque};

struct Solution;

impl Solution {
    pub fn nearest_exit(maze: Vec<Vec<char>>, entrance: Vec<i32>) -> i32 {
        const EMPTY_CELL: char = '.';
        const WALL_CELL: char = '+';

        let directions = [(1, 0), (-1, 0), (0, 1), (0, -1)];
        let (rows, cols) = (maze.len(), maze[0].len());
        let entrance_coords = [entrance[0] as usize, entrance[1] as usize];

        let mut queue = VecDeque::new();
        let mut visited = HashSet::new();
        let mut steps = 0;

        queue.push_back(entrance_coords);

        while !queue.is_empty() {
            let size = queue.len();

            for _ in 0..size {
                let coords = queue.pop_front().unwrap();

                if visited.contains(&coords) {
                    continue;
                }
                visited.insert(coords);

                let [i, j] = coords;

                if coords != entrance_coords && (i == 0 || i == rows - 1 || j == 0 || j == cols - 1)
                {
                    return steps;
                }

                for (dx, dy) in directions {
                    let new_i = i as i32 + dy;
                    let new_j = j as i32 + dx;

                    if new_j < 0 || new_j >= cols as i32 || new_i < 0 || new_i >= rows as i32 {
                        continue;
                    }

                    let new_i = new_i as usize;
                    let new_j = new_j as usize;

                    if maze[new_i][new_j] == EMPTY_CELL && !visited.contains(&[new_i, new_j]) {
                        queue.push_back([new_i, new_j])
                    }
                }
            }
            steps += 1;
        }

        -1
    }
}
