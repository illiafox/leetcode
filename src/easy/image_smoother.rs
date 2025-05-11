struct Solution;

impl Solution {
    pub fn image_smoother(img: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let m = img.len();
        let n = img[0].len();
        let mut smoothed = vec![vec![0; n]; m];

        for (i, row) in img.iter().enumerate() {
            for (j, _) in row.iter().enumerate() {
                let mut sum = 0;
                let mut count = 0;

                for di in -1..=1 {
                    for dj in -1..=1 {
                        let ni = i as isize + di;
                        let nj = j as isize + dj;
                        if ni >= 0 && ni < m as isize && nj >= 0 && nj < n as isize {
                            sum += img[ni as usize][nj as usize];
                            count += 1;
                        }
                    }
                }

                smoothed[i][j] = sum / count;
            }
        }

        smoothed
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            vec![vec![1, 1, 1], vec![1, 0, 1], vec![1, 1, 1]],
            vec![vec![0, 0, 0], vec![0, 0, 0], vec![0, 0, 0]],
        ),
        (
            vec![vec![100, 200, 100], vec![200, 50, 200], vec![100, 200, 100]],
            vec![
                vec![137, 141, 137],
                vec![141, 138, 141],
                vec![137, 141, 137],
            ],
        ),
    ];

    for (img, expected) in test_cases {
        assert_eq!(
            Solution::image_smoother(img.clone()),
            expected,
            "failed for img {img:?}",
        );
    }
}
