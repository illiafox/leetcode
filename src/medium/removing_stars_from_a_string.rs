struct Solution;

impl Solution {
    pub fn remove_stars(s: String) -> String {
        let mut stack = vec![];

        for c in s.bytes() {
            if c != b'*' {
                stack.push(c);
                continue;
            }

            while let Some(x) = stack.pop() {
                if x != b'*' {
                    break;
                }
            }
        }

        String::from_utf8(stack).unwrap()
    }
}

#[test]
fn test() {
    let test_cases = [("leet**cod*e", "lecoe"), ("erase*****", "")];

    for &(input, expected) in &test_cases {
        assert_eq!(
            Solution::remove_stars(input.to_string()),
            expected,
            "failed for input: {}",
            expected
        );
    }
}
