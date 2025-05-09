struct Solution;

use std::cmp::Ordering;
use std::collections::{BinaryHeap, HashMap};

#[derive(Debug, PartialEq)]
struct State {
    number: i32,
    occurrences: i32,
}

impl Eq for State {}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.occurrences.cmp(&other.occurrences))
    }
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        other.partial_cmp(self).unwrap()
    }
}

impl Solution {
    pub fn rearrange_barcodes(barcodes: Vec<i32>) -> Vec<i32> {
        let mut counter: HashMap<i32, i32> = HashMap::new();
        for &x in &barcodes {
            *counter.entry(x).or_insert(0) += 1;
        }

        let mut out = Vec::with_capacity(barcodes.len());
        let mut heap = BinaryHeap::with_capacity(counter.len());

        for (number, occurrences) in counter {
            heap.push(State {
                number,
                occurrences,
            });
        }

        let mut prev: Option<State> = None;

        while let Some(mut b) = heap.pop().or_else(|| prev.take()) {
            out.push(b.number);
            b.occurrences -= 1;
            if let Some(old) = prev.take() {
                heap.push(old);
            }
            if b.occurrences > 0 {
                prev = Some(b);
            }
        }

        out
    }
}

#[test]
fn test() {
    fn is_valid_rearrangement(barcodes: &[i32]) -> bool {
        for i in 1..barcodes.len() {
            if barcodes[i] == barcodes[i - 1] {
                return false;
            }
        }
        true
    }

    let test_cases = [
        vec![1, 1, 1, 2, 2, 2],
        vec![1, 1, 1, 1, 2, 2, 3, 3],
        vec![2, 2, 2, 1, 5],
    ];

    for barcodes in test_cases {
        let output = Solution::rearrange_barcodes(barcodes.clone());

        assert!(
            is_valid_rearrangement(&output),
            "failed for barcode: {:?} output: {:?}",
            barcodes,
            output
        );
    }
}
