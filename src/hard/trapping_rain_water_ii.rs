struct Solution;

use std::cmp::{max, Ordering};
use std::collections::BinaryHeap;

#[derive(Eq, PartialEq)]
struct Cell {
    coords: (usize, usize),
    height: i32,
}

impl Ord for Cell {
    fn cmp(&self, other: &Self) -> Ordering {
        other.height.cmp(&self.height)
    }
}

impl PartialOrd for Cell {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Solution {
    fn new_cell(coords: (usize, usize), height: i32) -> Cell {
        Cell { coords, height }
    }

    pub fn trap_rain_water(height_map: Vec<Vec<i32>>) -> i32 {
        let directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]; // Possible moves.

        let rows = height_map.len();
        let cols = height_map[0].len();

        let mut visited = vec![vec![false; cols]; rows];

        let mut boundary: BinaryHeap<Cell> = BinaryHeap::new();

        for i in 0..rows {
            boundary.push(Self::new_cell((i, 0), height_map[i][0]));
            boundary.push(Self::new_cell((i, cols - 1), height_map[i][cols - 1]));

            visited[i][0] = true;
            visited[i][cols - 1] = true;
        }

        for i in 0..cols {
            boundary.push(Self::new_cell((0, i), height_map[0][i]));
            boundary.push(Self::new_cell((rows - 1, i), height_map[rows - 1][i]));

            visited[0][i] = true;
            visited[rows - 1][i] = true;
        }

        let mut total_water_volume = 0;

        while let Some(cell) = boundary.pop() {
            let (i, j) = cell.coords;
            let min_boundary_height = cell.height;

            for (dx, dy) in directions {
                let new_i = i as i32 + dy;
                let new_j = j as i32 + dx;

                if new_j < 0 || new_j >= cols as i32 || new_i < 0 || new_i >= rows as i32 {
                    continue;
                }

                let new_i = new_i as usize;
                let new_j = new_j as usize;

                if visited[new_i][new_j] {
                    continue;
                }

                let neighbour_height = height_map[new_i][new_j];
                if neighbour_height < min_boundary_height {
                    total_water_volume += min_boundary_height - neighbour_height
                }

                boundary.push(Self::new_cell(
                    (new_i, new_j),
                    max(neighbour_height, min_boundary_height),
                ));
                visited[new_i][new_j] = true
            }
        }

        total_water_volume
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            vec![
                vec![1, 4, 3, 1, 3, 2],
                vec![3, 2, 1, 3, 2, 4],
                vec![2, 3, 3, 2, 3, 1],
            ],
            4,
        ),
        (
            vec![
                vec![3, 3, 3, 3, 3],
                vec![3, 2, 2, 2, 3],
                vec![3, 2, 1, 2, 3],
                vec![3, 2, 2, 2, 3],
                vec![3, 3, 3, 3, 3],
            ],
            10,
        ),
        (
            vec![
                vec![12, 13, 1, 12],
                vec![13, 4, 13, 12],
                vec![13, 8, 10, 12],
                vec![12, 13, 12, 12],
                vec![13, 13, 13, 13],
            ],
            14,
        ),
    ];

    for (a, expected) in test_cases {
        assert_eq!(
            Solution::trap_rain_water(a.clone()),
            expected,
            "failed for input: height_map: {:?}",
            a
        );
    }
}
