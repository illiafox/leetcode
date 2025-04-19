struct Solution;

impl Solution {
    fn lower_bound(nums: &[i32], mut low: usize, mut high: usize, key: i32) -> usize {
        while low < high {
            let mid = low + (high - low) / 2;
            if nums[mid] < key {
                low = mid + 1
            } else {
                high = mid
            }
        }

        low
    }

    pub fn count_fair_pairs(nums: Vec<i32>, lower: i32, upper: i32) -> i64 {
        let mut nums = nums;
        nums.sort();
        let mut pairs: i64 = 0;

        for (i, n) in nums.iter().enumerate() {
            let res_lower = Self::lower_bound(&nums, i + 1, nums.len(), lower - n);
            let res_higher = Self::lower_bound(&nums, i + 1, nums.len(), upper - n + 1);

            pairs += (res_higher - res_lower) as i64;
        }

        pairs
    }
}
