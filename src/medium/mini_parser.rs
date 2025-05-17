struct Solution;

#[derive(Debug, PartialEq, Eq)]
enum NestedInteger {
    Int(i32),
    List(Vec<NestedInteger>),
}

impl Solution {
    // wrapper for leetcode
    pub fn deserialize(s: String) -> NestedInteger {
        Self::deserialize_str(&s)
    }

    pub fn deserialize_str(s: &str) -> NestedInteger {
        if s.is_empty() {
            return NestedInteger::List(vec![]);
        }

        let chars = s.as_bytes();

        // base case: single integer
        if chars[0] != b'[' {
            return NestedInteger::Int(s.parse::<i32>().unwrap());
        }

        // nested list
        let mut result = vec![];
        let mut start = 1; // skip initial '['
        let mut depth = 0;

        for i in 1..s.len() - 1 {
            match chars[i] {
                b',' if depth == 0 => {
                    result.push(Self::deserialize_str(&s[start..i]));
                    start = i + 1;
                }
                b'[' => depth += 1,
                b']' => depth -= 1,
                _ => {}
            }
        }

        // push last element
        if start < s.len() - 1 {
            result.push(Self::deserialize_str(&s[start..s.len() - 1]));
        }

        NestedInteger::List(result)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use NestedInteger::{Int, List};

    #[test]
    fn test_cases() {
        // https://leetcode.com/problems/mini-parser/

        let testcases: Vec<(&str, NestedInteger)> = vec![
            ("324", Int(324)),
            (
                "[123,[456,[789]]]",
                List(vec![Int(123), List(vec![Int(456), List(vec![Int(789)])])]),
            ),
            (
                "[123,456,[788,799,833],[[]],10,[]]",
                List(vec![
                    Int(123),
                    Int(456),
                    List(vec![Int(788), Int(799), Int(833)]),
                    List(vec![List(vec![])]),
                    Int(10),
                    List(vec![]),
                ]),
            ),
        ];

        for (input, expected) in testcases {
            assert_eq!(
                Solution::deserialize(input.to_string()),
                expected,
                "failed for input {input}",
            );
        }
    }
}
