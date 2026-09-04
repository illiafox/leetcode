struct Solution;

// https://leetcode.com/problems/smallest-stable-index-i
impl Solution {
    pub fn first_stable_index(nums: Vec<i32>, k: i32) -> i32 {
        if nums.is_empty() {
            return -1;
        }

        // (value, index)
        let mut min_stack: Vec<(i32, usize)> = Vec::new();

        for (i, &val) in nums.iter().enumerate().rev() {
            if min_stack.last().is_none_or(|&(last_min, _)| val < last_min) {
                min_stack.push((val, i));
            }
        }

        let mut running_max = nums[0];

        for (i, &n) in nums.iter().enumerate() {
            running_max = running_max.max(n);

            let min_n = min_stack.last().unwrap().0;

            if running_max - min_n <= k {
                return i as i32;
            }

            if min_stack.last().unwrap().1 == i {
                min_stack.pop();
            }
        }

        -1
    }

    pub fn first_stable_index_first_try(nums: Vec<i32>, k: i32) -> i32 {
        let l = nums.len();

        let mut min_n = vec![0; l];
        min_n[l - 1] = nums[l - 1];

        let mut max_n = vec![0; l];
        max_n[0] = nums[0];

        for i in 1..l {
            max_n[i] = i32::max(max_n[i - 1], nums[i]);
            min_n[l - i - 1] = i32::min(min_n[l - i], nums[l - i - 1]);
        }

        for i in 0..l {
            if max_n[i] - min_n[i] <= k {
                return i as i32;
            }
        }

        -1
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    type SolutionFn = fn(Vec<i32>, i32) -> i32;

    #[test]
    fn test_all_implementations() {
        let solutions: &[(&str, SolutionFn)] = &[
            ("first_stable_index", Solution::first_stable_index),
            (
                "first_stable_index_first_try",
                Solution::first_stable_index_first_try,
            ),
        ];

        let test_cases = [
            (vec![5, 0, 1, 4], 3, 3),
            (vec![3, 2, 1], 1, -1),
            (vec![0], 0, 0),
            (vec![0, 0], 0, 0),
            (vec![6, 1, 4], 5, 0),
        ];

        for (fn_name, func) in solutions {
            for (idx, (nums, k, expected)) in test_cases.iter().enumerate() {
                let actual = func(nums.clone(), *k);
                assert_eq!(
                    actual, *expected,
                    "[{fn_name}] Failed case #{idx}: nums = {nums:?}, k = {k} (expected {expected}, got {actual})"
                );
            }
        }
    }
}
