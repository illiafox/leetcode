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

    pub fn first_stable_index_sec_try(nums: Vec<i32>, k: i32) -> i32 {
        let mut min_prefix = Vec::with_capacity(nums.len());

        for (i, &n) in nums.iter().rev().enumerate() {
            if let Some(&prev) = min_prefix.last() {
                if prev == n && nums.get(nums.len() - i).is_some_and(|&x| n != x) {
                    min_prefix.push(n);
                }
                if prev > n {
                    min_prefix.push(n);
                }
                if prev < n {
                    min_prefix.push(prev);
                }
            } else {
                min_prefix.push(n);
            }
        }

        let mut max_n = nums[0];

        for (i, &n) in nums.iter().enumerate() {
            max_n = i32::max(max_n, n);

            let min_n = *min_prefix.last().expect("bad stuff happened");

            if max_n - min_n <= k {
                return i as i32;
            }

            if n != min_n || nums.get(i + 1).is_some_and(|&x| x != n) {
                min_prefix.pop();
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

        let mut min_index = -1;
        let mut min_stable = None;

        for i in 0..l {
            let stable = max_n[i] - min_n[i];
            if stable <= k {
                if let Some(prev) = min_stable
                    && prev < stable
                {
                    continue;
                }
                min_index = i as i32;
                min_stable = Some(stable);
            }
        }

        min_index
    }
}
