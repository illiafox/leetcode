struct Solution;

impl Solution {
    pub fn num_decodings(s: String) -> i32 {
        let bytes = s.as_bytes();
        let n = s.len();

        if n == 0 || bytes[0] == b'0' {
            return 0;
        }

        let mut dp = vec![0; n + 1];
        dp[0] = 1;
        dp[1] = 1;

        for i in 2..=n {
            let one = bytes[i - 1];
            let two = (bytes[i - 2] - b'0') * 10 + (bytes[i - 1] - b'0');

            if one != b'0' {
                dp[i] += dp[i - 1];
            }

            if (10..=26).contains(&two) {
                dp[i] += dp[i - 2];
            }
        }

        dp[n]
    }
}
