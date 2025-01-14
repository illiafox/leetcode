use std::collections::HashMap;
struct Solution;

impl Solution {
    pub fn find_the_prefix_common_array(a: Vec<i32>, b: Vec<i32>) -> Vec<i32> {
        let mut frequencies: Vec<i32> = vec![0; a.len()];

        let mut answer: Vec<i32> = vec![0; a.len()];
        let mut common = 0;

        for i in 0..a.len() {
            let a_idx = a[i] as usize - 1;
            frequencies[a_idx] += 1;
            if frequencies[a_idx] == 2 {
                common += 1
            }

            let b_idx = b[i] as usize - 1;
            frequencies[b_idx] += 1;
            if frequencies[b_idx] == 2 {
                common += 1
            }

            answer[i] = common
        }

        answer
    }
    pub fn find_the_prefix_common_array_slow(a: Vec<i32>, b: Vec<i32>) -> Vec<i32> {
        let map: HashMap<i32, usize> = a
            .iter()
            .enumerate()
            .map(|(index, &value)| (value, index))
            .collect();

        let mut answer: Vec<i32> = vec![0; a.len()];

        for (i, answ) in answer.iter_mut().enumerate() {
            let mut common = 0;

            for v in b.iter().take(i + 1) {
                if let Some(f) = map.get(v) {
                    if *f <= i {
                        common += 1
                    }
                }
            }

            *answ = common
        }

        answer
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![1, 3, 2, 4], vec![3, 1, 2, 4], vec![0, 2, 3, 4]),
        (vec![2, 3, 1], vec![3, 1, 2], vec![0, 1, 3]),
    ];

    for (a, b, expected) in test_cases {
        assert_eq!(
            Solution::find_the_prefix_common_array_slow(a.clone(), b.clone()),
            expected,
            "find_the_prefix_common_array_slow: failed for a: {:?} b: {:?}",
            a,
            b
        );

        assert_eq!(
            Solution::find_the_prefix_common_array(a.clone(), b.clone()),
            expected,
            "find_the_prefix_common_array: failed for a: {:?} b: {:?}",
            a,
            b
        );
    }
}
