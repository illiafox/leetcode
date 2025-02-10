struct Solution;

use std::collections::HashSet;

impl Solution {
    pub fn clear_digits(s: String) -> String {
        let bytes = s.as_bytes();
        let mut output: Vec<u8> = Vec::new();
        let mut marked_indices: HashSet<usize> = HashSet::new();

        for (i, &ch) in bytes.iter().rev().enumerate() {
            if ch.is_ascii_digit() {
                for j in (0..s.len() - i).rev() {
                    if bytes[j].is_ascii_lowercase() && !marked_indices.contains(&j) {
                        marked_indices.insert(j);
                        break;
                    }
                }
            }
        }

        for (i, &ch) in bytes.iter().enumerate() {
            if ch.is_ascii_digit() {
                continue;
            }
            if !marked_indices.contains(&i) {
                output.push(ch);
            }
        }

        String::from_utf8(output).unwrap()
    }
}

#[test]
fn test() {
    let test_cases = [("abc", "abc"), ("cb34", ""), ("a8f", "f")];

    for (input, expected) in test_cases {
        assert_eq!(
            Solution::clear_digits(input.to_string()),
            expected,
            "failed for input {:?}",
            input
        );
    }
}
