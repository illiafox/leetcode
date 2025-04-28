use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut frequencies: HashMap<i32, i32> = HashMap::new();
        let mut unique_elems = vec![];

        for n in nums {
            if *frequencies
                .entry(n)
                .and_modify(|counter| *counter += 1)
                .or_insert(1)
                == 1
            {
                unique_elems.push(n)
            }
        }

        unique_elems.sort_by(|a, b| frequencies.get(b).unwrap().cmp(frequencies.get(a).unwrap()));
        unique_elems.truncate(k as usize);

        unique_elems
    }
}
