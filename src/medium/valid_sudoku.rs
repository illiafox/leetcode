struct Solution;

impl Solution {
    // https://leetcode.com/problems/valid-sudoku
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut rows = [0u16; 9];
        let mut cols = [0u16; 9];
        let mut boxes = [0u16; 9];

        for r in 0..9 {
            for (c, &ch) in board[r].iter().enumerate() {
                if ch == '.' {
                    continue;
                }
                if !('1'..='9').contains(&ch) {
                    return false;
                }

                let bit = 1u16 << ((ch as u8 - b'1') as u16);
                let b = (r / 3) * 3 + (c / 3);

                if (rows[r] & bit) != 0 || (cols[c] & bit) != 0 || (boxes[b] & bit) != 0 {
                    return false;
                }

                rows[r] |= bit;
                cols[c] |= bit;
                boxes[b] |= bit;
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
