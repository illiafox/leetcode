use std::collections::VecDeque;

struct Solution;

#[derive(Clone, Copy)]
struct Task {
    i: u8,
    j: u8,
    energy: i8,
    mask: u16,
}

// https://leetcode.com/problems/minimum-moves-to-clean-the-classroom
//
// Optimize for the following constraints:
// 1 <= m == classroom.length <= 20
// 1 <= n == classroom[i].length <= 20
// classroom[i][j] is one of 'S', 'L', 'R', 'X', or '.'
// 1 <= energy <= 50
// There is exactly one 'S' in the grid.
// There are at most 10 'L' cells in the grid.
impl Solution {
    pub fn min_moves(classroom: Vec<String>, max_energy: i32) -> i32 {
        let rows = classroom.len();
        let cols = classroom[0].len();
        let grid: Vec<&[u8]> = classroom.iter().map(|s| s.as_bytes()).collect();

        let mut start = None;
        let mut litter_id = [[None; 20]; 20];
        let mut total_litter: u8 = 0;

        for r in 0..rows {
            for c in 0..cols {
                match grid[r][c] {
                    b'S' => start = Some((r as u8, c as u8)),
                    b'L' => {
                        litter_id[r][c] = Some(total_litter);
                        total_litter += 1;
                    }
                    _ => {}
                }
            }
        }

        if total_litter == 0 {
            return 0;
        }

        let (sr, sc) = start.expect("start not found");
        let target_mask = (1u16 << total_litter) - 1;
        let num_cells = rows * cols;
        let max_energy = max_energy as i8;

        // Optimization:  we convert a 3D array of [i][j][mask] to the 1D array
        let total_states = (1usize << total_litter) * num_cells;
        let mut best_energy = vec![-1; total_states];

        // in the BFS we mostly check adjacent cells
        //
        // if we pick the formula "idx = (i * rows + j) * num_cells * mask",
        // for example in 32*32=1024 grid seeking from (0, 0) to (1, 0)
        // will take moving up to 20480 elements!!
        //
        // that's why we use inverted indexing: as mask does not change often, we use it first
        let get_idx = |r: u8, c: u8, mask: u16| -> usize {
            (mask as usize) * num_cells + (r as usize * cols + c as usize)
        };

        best_energy[get_idx(sr, sc, 0)] = max_energy;

        let mut queue = VecDeque::with_capacity(1024);
        queue.push_back(Task {
            i: sr,
            j: sc,
            energy: max_energy,
            mask: 0,
        });

        let directions: [(i8, i8); 4] = [(1, 0), (-1, 0), (0, 1), (0, -1)];
        let mut step = 0;

        while !queue.is_empty() {
            let level_size = queue.len();

            for _ in 0..level_size {
                let curr = queue.pop_front().unwrap();

                if curr.mask == target_mask {
                    return step;
                }

                if curr.energy <= 0 {
                    continue;
                }

                for (dr, dc) in directions {
                    let nr = curr.i as i8 + dr;
                    let nc = curr.j as i8 + dc;

                    if nr < 0 || nr >= rows as i8 || nc < 0 || nc >= cols as i8 {
                        continue;
                    }

                    let (nr, nc) = (nr as u8, nc as u8);
                    let cell = grid[nr as usize][nc as usize];

                    if cell == b'X' {
                        continue;
                    }

                    let next_energy = if cell == b'R' {
                        max_energy
                    } else {
                        curr.energy - 1
                    };

                    let next_mask = match litter_id[nr as usize][nc as usize] {
                        Some(id) => curr.mask | (1 << id),
                        None => curr.mask,
                    };

                    let idx = get_idx(nr, nc, next_mask);
                    if next_energy <= best_energy[idx] {
                        continue;
                    }

                    best_energy[idx] = next_energy;
                    queue.push_back(Task {
                        i: nr,
                        j: nc,
                        energy: next_energy,
                        mask: next_mask,
                    });
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
