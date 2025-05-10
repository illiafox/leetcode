struct Solution;

impl Solution {
    pub fn num_tilings(n: i32) -> i32 {
        match n {
            0 => return 1,
            1 => return 1,
            2 => return 2,
            _ => {}
        }

        let mut dp = vec![0; (n + 1) as usize];
        dp[0] = 1;
        dp[1] = 1;
        dp[2] = 2;

        // "Since the answer may be very large, return it modulo 10^9 + 7"
        const MOD: i64 = 1_000_000_007;
        for i in 3..=n as usize {
            dp[i] = (2 * dp[i - 1] + dp[i - 3]) % MOD
        }

        dp[n as usize] as i32
    }
}

#[test]
fn test() {
    let test_cases = [(3, 5), (1, 1)];

    for (n, expected) in test_cases {
        assert_eq!(Solution::num_tilings(n), expected, "failed for n = {n}");
    }
}
