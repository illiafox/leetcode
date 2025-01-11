struct Solution;

impl Solution {
    const ENGLISH_LETTERS_COUNT: usize = 26;

    pub fn can_construct(s: String, k: i32) -> bool {
        if s.len() < k as usize {
            return false;
        }

        let chars = Self::count_chars(s);

        let odd_chars_count = chars.iter().filter(|&&num| num % 2 == 1).count();

        if odd_chars_count > k as usize {
            return false;
        }

        true
    }

    fn count_chars(s: String) -> [usize; Self::ENGLISH_LETTERS_COUNT] {
        let mut chars: [usize; Self::ENGLISH_LETTERS_COUNT] = Default::default();

        for ch in s.chars() {
            chars[ch as usize - 'a' as usize] += 1;
        }

        chars
    }
}

#[test]
fn test() {
    let test_cases = [
        ("annabelle", 2, true),
        ("leetcode", 3, false),
        ("true", 4, true),
        ("yzyzyzyzyzyzyzy", 2, true),
    ];

    for &(input, k, expected) in &test_cases {
        assert_eq!(
            Solution::can_construct(input.to_string(), k),
            expected,
            "failed for input: {}, k: {}",
            input,
            k
        );
    }
}
