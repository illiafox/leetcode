use std::collections::HashSet;

struct Solution;

impl Solution {
    pub fn tuple_same_product(nums: Vec<i32>) -> i32 {
        let nums_len = nums.len();
        let mut nums = nums.clone();
        nums.sort_unstable();

        let mut total = 0;

        for a_index in 0..nums_len {
            for b_index in (a_index + 1..nums_len).rev() {
                let product = nums[a_index] * nums[b_index];

                let mut possible_values: HashSet<i32> = HashSet::new();

                for &n in nums.iter().take(b_index).skip(a_index + 1) {
                    if product % n == 0 {
                        let d = product / n;

                        if possible_values.contains(&d) {
                            total += 8;
                        }

                        possible_values.insert(n);
                    }
                }
            }
        }

        total
    }
}
