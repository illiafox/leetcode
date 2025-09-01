struct Solution;

use std::cmp::Ordering;
use std::collections::BinaryHeap;

#[derive(Debug, PartialEq)]
struct Class {
    total: i32,
    pass: i32,
    gain: f64,
}

impl Eq for Class {}

impl Ord for Class {
    fn cmp(&self, other: &Self) -> Ordering {
        self.gain.partial_cmp(&other.gain).unwrap()
    }
}

impl PartialOrd for Class {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Solution {
    #[inline(always)]
    fn calculate_cain(total: i32, pass: i32) -> f64 {
        ((pass as f64 + 1.0) / (total as f64 + 1.0)) - (pass as f64 / total as f64)
    }

    // https://leetcode.com/problems/maximum-average-pass-ratio
    pub fn max_average_ratio(classes: Vec<Vec<i32>>, extra_students: i32) -> f64 {
        let mut heap: BinaryHeap<Class> = classes
            .iter()
            .map(|v| -> Class {
                let total = v[1];
                let pass = v[0];
                let gain = Self::calculate_cain(total, pass);
                Class { total, pass, gain }
            })
            .collect();

        for _ in 0..extra_students {
            let mut m = heap.pop().unwrap();
            m.pass += 1;
            m.total += 1;
            m.gain = Self::calculate_cain(m.total, m.pass);
            heap.push(m);
        }

        let mut ratio_sum = 0.0;

        while let Some(m) = heap.pop() {
            ratio_sum += m.pass as f64 / m.total as f64
        }

        ratio_sum / classes.len() as f64
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            vec![vec![2, 4], vec![3, 9], vec![4, 5], vec![2, 10]],
            4,
            0.53485,
        ),
        (vec![vec![1, 2], vec![3, 5], vec![2, 2]], 2, 0.78333),
    ];

    for (classes, extra_students, expected) in test_cases {
        let got = Solution::max_average_ratio(classes.clone(), extra_students);
        assert!(
            // answers within 10-5 of the actual answer are accepted.
            (got - expected).abs() < 1e-5,
            "failed for {classes:?}: got {got}, expected {expected}"
        );
    }
}
