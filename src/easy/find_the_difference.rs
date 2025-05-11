struct Solution;

impl Solution {
    pub fn find_the_difference(s: String, t: String) -> char {
        let s_chars = Self::count_chars(s);
        let t_chars = Self::count_chars(t);

        for i in 0..s_chars.len() {
            if s_chars[i] != t_chars[i] {
                return (i as u8 + b'a') as char;
            }
        }

        unreachable!()
    }

    const ENGLISH_LETTERS_COUNT: usize = 26;
    fn count_chars(s: String) -> [usize; Self::ENGLISH_LETTERS_COUNT] {
        let mut chars: [usize; Self::ENGLISH_LETTERS_COUNT] = Default::default();

        for ch in s.chars() {
            chars[ch as usize - 'a' as usize] += 1;
        }

        chars
    }
}
