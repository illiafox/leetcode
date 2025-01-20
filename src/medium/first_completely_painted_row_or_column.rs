use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn first_complete_index(arr: Vec<i32>, mat: Vec<Vec<i32>>) -> i32 {
        let rows = mat.len();
        let cols = mat[0].len();

        let mut row_count = vec![0; rows];
        let mut col_count = vec![0; cols];

        let mut value_to_position = HashMap::new();

        for (i, row) in mat.iter().enumerate() {
            for (j, &value) in row.iter().enumerate() {
                value_to_position.insert(value, (i, j));
            }
        }

        for (idx, value) in arr.iter().enumerate() {
            if let Some(&(i, j)) = value_to_position.get(value) {
                row_count[i] += 1;
                if row_count[i] == cols {
                    return idx as i32;
                }

                col_count[j] += 1;
                if col_count[j] == rows {
                    return idx as i32;
                }
            }
        }

        -1
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![1, 3, 4, 2], vec![vec![1, 4], vec![2, 3]], 2),
        (
            vec![2, 8, 7, 4, 1, 3, 5, 6, 9],
            vec![vec![3, 2, 5], vec![1, 4, 6], vec![8, 7, 9]],
            3,
        ),
        (
            vec![1, 4, 5, 2, 6, 3],
            vec![vec![4, 3, 5], vec![1, 2, 6]],
            1,
        ),
    ];

    for (arr, mat, expected) in test_cases {
        assert_eq!(
            Solution::first_complete_index(arr.clone(), mat.clone()),
            expected,
            "failed for arr: {:?}, mat: {:?}",
            arr,
            mat
        );
    }
}
