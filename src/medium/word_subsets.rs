struct Solution;

impl Solution {
    const ENGLISH_LETTERS_COUNT: usize = 26;

    pub fn word_subsets(words1: Vec<String>, words2: Vec<String>) -> Vec<String> {
        let mut subset_requirements: [usize; Self::ENGLISH_LETTERS_COUNT] = Default::default();

        for word in words2 {
            let chars = Self::count_chars(&word);

            for (ch, &count) in chars.iter().enumerate() {
                if subset_requirements[ch] < count {
                    subset_requirements[ch] = count
                }
            }
        }

        let mut out: Vec<String> = vec![];

        'outer: for word in words1 {
            let chars = Self::count_chars(&word);

            for (ch, &count) in chars.iter().enumerate() {
                if count < subset_requirements[ch] {
                    continue 'outer;
                }
            }

            out.push(word);
        }

        out
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
    assert_eq!(
        Solution::word_subsets(
            vec![
                "amazon".to_string(),
                "apple".to_string(),
                "facebook".to_string(),
                "google".to_string(),
                "leetcode".to_string(),
            ],
            vec!["e".to_string(), "o".to_string()]
        ),
        vec![
            "facebook".to_string(),
            "google".to_string(),
            "leetcode".to_string(),
        ]
    );

    assert_eq!(
        Solution::word_subsets(
            vec![
                "amazon".to_string(),
                "apple".to_string(),
                "facebook".to_string(),
                "google".to_string(),
                "leetcode".to_string(),
            ],
            vec!["l".to_string(), "e".to_string()]
        ),
        vec![
            "apple".to_string(),
            "google".to_string(),
            "leetcode".to_string(),
        ]
    );
}
