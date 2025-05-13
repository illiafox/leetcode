struct Solution;

impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut buf = vec![0; 9];

        let is_valid = |buf: &[u8]| buf.iter().all(|&x| x <= 1);
        let update = |buf: &mut Vec<u8>, x: char| {
            if x != '.' {
                buf[(x as u8 - b'1') as usize] += 1
            }
        };

        for i in 0..board.len() {
            buf.fill(0);
            for j in 0..board[i].len() {
                update(&mut buf, board[i][j])
            }
            if !is_valid(&buf) {
                return false;
            }
        }
        for j in 0..board[0].len() {
            buf.fill(0);
            for i in 0..board.len() {
                update(&mut buf, board[i][j])
            }
            if !is_valid(&buf) {
                return false;
            }
        }

        let offsets = [
            (0, 0),
            (3, 0),
            (6, 0),
            (0, 3),
            (3, 3),
            (6, 3),
            (0, 6),
            (3, 6),
            (6, 6),
        ];

        for (dx, dy) in offsets {
            buf.fill(0);

            for i in dy..dy + 3 {
                for j in dx..dx + 3 {
                    update(&mut buf, board[i][j])
                }
            }
            if !is_valid(&buf) {
                return false;
            }
        }

        true
    }
}
