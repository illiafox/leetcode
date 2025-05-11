struct Solution;

impl Solution {
    const ENGLISH_LETTERS_COUNT: usize = 26;
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut chars: [i32; Self::ENGLISH_LETTERS_COUNT] = Default::default();

        for ch in s.chars() {
            chars[ch as usize - 'a' as usize] += 1;
        }
        for ch in t.chars() {
            chars[ch as usize - 'a' as usize] -= 1;
        }

        chars.iter().all(|&x| x == 0)
    }
}

#[test]
fn test() {
    let test_cases = [("anagram", "nagaram", true), ("rat", "car", false)];

    for (s, t, expected) in test_cases {
        assert_eq!(
            Solution::is_anagram(s.to_string(), t.to_string()),
            expected,
            "failed for s = {s} t = {t}",
        );
    }
}
