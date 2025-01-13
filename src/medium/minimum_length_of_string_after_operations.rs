struct Solution;

impl Solution {
    const ENGLISH_LETTERS_COUNT: usize = 26;

    /// https://leetcode.com/problems/minimum-length-of-string-after-operations/
    pub fn minimum_length(s: String) -> i32 {
        let chars = Self::count_chars(&s);

        let mut final_length = s.len();

        for &count in chars.iter() {
            let mut letters_left = count;

            while letters_left > 2 {
                let v = letters_left / 3 * 2;
                final_length -= v;
                letters_left -= v;
            }
        }

        final_length as i32
    }

    fn count_chars(s: &str) -> [usize; Self::ENGLISH_LETTERS_COUNT] {
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
        ("abaacbcbb", 5),
        ("aa", 2),
        ("ucvbutgkohgbcobqeyqwppbxqoynxeuuzouyvmydfhrprdbuzwqebwuiejoxsxdhbmuaiscalnteocghnlisxxawxgcjloevrdcj", 38),
    ];

    for &(input, expected) in &test_cases {
        assert_eq!(
            Solution::minimum_length(input.to_string()),
            expected,
            "failed for input: {}, k: {}",
            input,
            expected
        );
    }
}
