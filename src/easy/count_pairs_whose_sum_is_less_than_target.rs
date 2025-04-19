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

    pub fn count_pairs(nums: Vec<i32>, target: i32) -> i32 {
        let mut nums = nums;
        nums.sort();
        let mut pairs = 0;

        for (i, n) in nums.iter().enumerate() {
            let res = Self::lower_bound(&nums, i + 1, nums.len(), target - n);
            pairs += (res - i - 1) as i32;
        }

        pairs
    }
}
