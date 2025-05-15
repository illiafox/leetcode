struct Solution;

impl Solution {
    pub fn set_zeroes(matrix: &mut [Vec<i32>]) {
        let rows = matrix.len();
        let cols = matrix[0].len();
        let mut first_row_zero = false;
        let mut first_col_zero = false;

        // check if first row has zero
        for j in 0..cols {
            if matrix[0][j] == 0 {
                first_row_zero = true;
                break;
            }
        }

        // check if first column has zero
        for row in matrix.iter().take(rows) {
            if row[0] == 0 {
                first_col_zero = true;
                break;
            }
        }

        // use first row and column to mark zero rows/cols
        for i in 1..rows {
            for j in 1..cols {
                if matrix[i][j] == 0 {
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                }
            }
        }

        // set zeroes based on markers
        for i in 1..rows {
            for j in 1..cols {
                if matrix[i][0] == 0 || matrix[0][j] == 0 {
                    matrix[i][j] = 0;
                }
            }
        }

        // set zeroes in first row if needed
        if first_row_zero {
            for j in 0..cols {
                matrix[0][j] = 0;
            }
        }

        // set zeroes in first column if needed
        if first_col_zero {
            for col in matrix.iter_mut().take(rows) {
                col[0] = 0;
            }
        }
    }
}

#[test]
fn test() {
    // https://leetcode.com/problems/set-matrix-zeroes/

    let test_cases = [
        (
            vec![vec![1, 1, 1], vec![1, 0, 1], vec![1, 1, 1]],
            vec![vec![1, 0, 1], vec![0, 0, 0], vec![1, 0, 1]],
        ),
        (
            vec![vec![0, 1, 2, 0], vec![3, 4, 5, 2], vec![1, 3, 1, 5]],
            vec![vec![0, 0, 0, 0], vec![0, 4, 5, 0], vec![0, 3, 1, 0]],
        ),
    ];

    for (mut matrix, expected) in test_cases {
        Solution::set_zeroes(&mut matrix);

        assert_eq!(matrix, expected, "failed for matrix: {matrix:?}",);
    }
}
