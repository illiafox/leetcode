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

        for row in &board {
            buf.fill(0);
            for &c in row {
                update(&mut buf, c)
            }
            if !is_valid(&buf) {
                return false;
            }
        }
        for j in 0..board[0].len() {
            buf.fill(0);
            for row in &board {
                update(&mut buf, row[j])
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

            for row in board.iter().skip(dy).take(3) {
                for &c in row.iter().skip(dx).take(3) {
                    update(&mut buf, c)
                }
            }
            if !is_valid(&buf) {
                return false;
            }
        }

        true
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            vec![
                vec!['5', '3', '.', '.', '7', '.', '.', '.', '.'],
                vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
                vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
                vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
                vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
                vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
                vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
                vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
                vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
            ],
            true,
        ),
        (
            vec![
                vec!['8', '3', '.', '.', '7', '.', '.', '.', '.'],
                vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
                vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
                vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
                vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
                vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
                vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
                vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
                vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
            ],
            false,
        ),
    ];

    for (input, expected) in test_cases {
        assert_eq!(
            Solution::is_valid_sudoku(input.clone()),
            expected,
            "failed for input {input:?}",
        );
    }
}
