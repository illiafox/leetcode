struct Solution;

impl Solution {
    pub fn xor_all_nums(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let mut ans = 0;

        for &num in &nums1 {
            if nums2.len() % 2 == 1 {
                ans ^= num;
            }
        }

        for &num in &nums2 {
            if nums1.len() % 2 == 1 {
                ans ^= num;
            }
        }

        ans
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![1, 2], vec![3, 4], 0),
        (vec![2, 1, 3], vec![10, 2, 5, 0], 13),
    ];

    for (a, b, expected) in test_cases {
        assert_eq!(
            Solution::xor_all_nums(a.clone(), b.clone()),
            expected,
            "failed for input: a: {:?} b: {:?}",
            a,
            b
        );
    }
}
