use std::collections::HashMap;

struct Solution;

impl Solution {
    fn explore_island(
        island_id: i32,
        grid: &Vec<Vec<i32>>,
        visited: &mut Vec<Vec<i32>>,
        (i, j): (i32, i32),
    ) -> i32 {
        let directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]; // Possible moves.

        if j < 0 || j >= grid[0].len() as i32 || i < 0 || i >= grid.len() as i32 {
            return 0;
        }

        let i_c = i as usize;
        let j_c = j as usize;

        if visited[i_c][j_c] != -1 || grid[i_c][j_c] == 0 {
            return 0;
        }
        visited[i_c][j_c] = island_id;

        let mut area = 1;

        for (dx, dy) in directions {
            area += Self::explore_island(island_id, grid, visited, (i + dy, j + dx))
        }

        area
    }

    pub fn largest_island(grid: Vec<Vec<i32>>) -> i32 {
        let rows = grid.len();
        let cols = grid[0].len();

        let mut visited = vec![vec![-1; cols]; rows];

        let mut island_areas = HashMap::new();

        let mut island_id = 1;

        for (i, row) in grid.iter().enumerate() {
            for (j, &is_land) in row.iter().enumerate() {
                if is_land == 1 && visited[i][j] == -1 {
                    _ = island_areas.insert(
                        island_id,
                        Self::explore_island(island_id, &grid, &mut visited, (i as i32, j as i32)),
                    );
                    island_id += 1
                }
            }
        }

        let mut max_size = 0;

        for (i, row) in grid.iter().enumerate() {
            for (j, &is_land) in row.iter().enumerate() {
                if is_land == 0 {
                    let directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]; // Possible moves.

                    let mut temp_size = 1;

                    let mut used_islands = vec![false; island_id as usize];

                    for (dx, dy) in directions {
                        let new_i = i as i32 + dy;
                        let new_j = j as i32 + dx;

                        if new_j < 0 || new_j >= cols as i32 || new_i < 0 || new_i >= rows as i32 {
                            continue;
                        }

                        let new_i = new_i as usize;
                        let new_j = new_j as usize;

                        if visited[new_i][new_j] == -1 {
                            continue;
                        }

                        let cur_island_id = visited[new_i][new_j];
                        if used_islands[(cur_island_id - 1) as usize] {
                            continue;
                        }
                        used_islands[(cur_island_id - 1) as usize] = true;

                        temp_size += island_areas.get(&cur_island_id).unwrap();
                    }

                    if temp_size > max_size {
                        max_size = temp_size
                    }
                }
            }
        }

        if max_size == 0 && !island_areas.is_empty() {
            return island_areas.iter().fold(0, |acc, (_, area)| acc + area);
        }

        max_size
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![vec![1, 0], vec![0, 1]], 3),
        (vec![vec![1, 1], vec![1, 0]], 4),
        (vec![vec![1, 1], vec![1, 1]], 4),
        (vec![vec![0, 0], vec![1, 1]], 3),
    ];

    for (grid, expected) in test_cases {
        assert_eq!(
            Solution::largest_island(grid.clone()),
            expected,
            "failed for input: grid: {:?}",
            grid
        );
    }
}
