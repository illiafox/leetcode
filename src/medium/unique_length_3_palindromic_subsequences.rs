use std::collections::HashSet;

struct Solution;

// https://leetcode.com/problems/unique-length-3-palindromic-subsequences

impl Solution {
    /// Returns the number of unique palindromes of length three that are a subsequence of s. Works only for strings with lowercase english letters
    pub fn count_palindromic_subsequence(s: String) -> i32 {
        const ENGLISH_LETTERS_COUNT: usize = 26;

        let mut first: [Option<usize>; ENGLISH_LETTERS_COUNT] = Default::default();
        let mut last: [Option<usize>; ENGLISH_LETTERS_COUNT] = Default::default();

        let bytes = s.as_bytes();

        for (i, &c) in bytes.iter().enumerate() {
            let idx = (c - b'a') as usize;
            if first[idx].is_none() {
                first[idx] = Some(i);
            }
            last[idx] = Some(i);
        }

        let mut count = 0;

        for l in 0..first.len() {
            if let (Some(i), Some(j)) = (first[l], last[l]) {
                if i < j {
                    let unique_chars: HashSet<u8> = bytes[i + 1..j].iter().copied().collect();
                    count += unique_chars.len() as i32;
                }
            }
        }

        count
    }

    pub fn count_palindromic_subsequence_old(s: String) -> i32 {
        let bytes = s.as_bytes();

        let letters: HashSet<char> = s.chars().collect();

        let mut count: i32 = 0;

        for c in letters {
            if let (Some(i), Some(j)) = (s.find(c), s.rfind(c)) {
                if i < j {
                    let unique_chars: HashSet<&u8> = bytes[i + 1..j].iter().collect();
                    count += unique_chars.len() as i32;
                }
            }
        }

        count
    }
}

#[test]

fn test() {
    let test_cases = [("aabca", 3), ("adc", 0), ("bbcbaba", 4)];

    for (input, expected) in test_cases {
        assert_eq!(
            Solution::count_palindromic_subsequence(input.to_string()),
            expected,
            "count_palindromic_subsequence: failed for input: {}",
            input
        );

        assert_eq!(
            Solution::count_palindromic_subsequence_old(input.to_string()),
            expected,
            "count_palindromic_subsequence_old: failed for input: {}",
            input
        );
    }
}
