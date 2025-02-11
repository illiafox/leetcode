struct Solution;

impl Solution {
    pub fn remove_occurrences(s: String, part: String) -> String {
        let mut stack = vec![];

        for &s in s.as_bytes() {
            stack.push(s);

            if stack.len() >= part.len() && Self::check_match(&stack, part.as_bytes()) {
                for _ in 0..part.len() {
                    stack.pop();
                }
            }
        }

        String::from_utf8(stack).unwrap()
    }

    fn check_match(stack: &[u8], part: &[u8]) -> bool {
        &stack[stack.len() - part.len()..] == part
    }
}

#[test]

fn test() {
    let test_cases = [("daabcbaabcbc", "abc", "dab"), ("axxxxyyyyb", "xy", "ab")];

    for &(s, part, expected) in &test_cases {
        assert_eq!(
            Solution::remove_occurrences(s.to_string(), part.to_string()),
            expected,
            "failed for s: {}, part: {}",
            s,
            expected
        );
    }
}
