struct Solution;

// TODO: rewrite....
impl Solution {
    pub fn reverse_words(s: String) -> String {
        let mut out = String::new();

        let mut last = s.len() - 1;

        for (i, ch) in s.chars().rev().enumerate() {
            if ch != ' ' {
                continue;
            }

            let i = s.len() - i - 1;

            if last == i {
                if last == 0 {
                    return out;
                }

                last -= 1;
                continue;
            }

            if !out.is_empty() {
                out.push(' ')
            }
            out.push_str(&s[i + 1..=last]);

            if i == 0 {
                return out;
            }
            last = i - 1
        }

        if !out.is_empty() {
            out.push(' ')
        }
        out.push_str(&s[0..=last]);

        out
    }
}

#[test]
fn test() {
    let test_cases = [
        ("the sky is blue", "blue is sky the"),
        ("  hello world  ", "world hello"),
        ("a good   example", "example good a"),
        (" asdasd df f", "f df asdasd"),
    ];

    for &(input, expected) in &test_cases {
        assert_eq!(
            Solution::reverse_words(input.to_string()),
            expected,
            "failed for input: {}",
            expected
        );
    }
}
