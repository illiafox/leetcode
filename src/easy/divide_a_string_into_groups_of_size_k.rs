struct Solution;

impl Solution {
    pub fn divide_string(s: String, k: i32, fill: char) -> Vec<String> {
        let mut out = vec![];
        let k = k as usize;

        let mut i = 0;
        while i < s.len() {
            let mut chunk = s[i..s.len().min(i + k)].to_string();
            while chunk.len() < k {
                chunk.push(fill);
            }
            out.push(chunk);
            i += k;
        }

        out
    }
}

#[test]
fn test() {
    let test_cases = [
        ("abcdefghi", 3, 'x', vec!["abc", "def", "ghi"]),
        ("abcdefghij", 3, 'x', vec!["abc", "def", "ghi", "jxx"]),
    ];

    for (s, k, fill, expected) in test_cases {
        assert_eq!(
            Solution::divide_string(s.to_string(), k, fill),
            expected,
            "failed for {s}, k {k}, fill {fill}",
        );
    }
}
