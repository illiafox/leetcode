struct Solution;

impl Solution {
    pub fn compress(chars: &mut [char]) -> i32 {
        let mut write_idx = 0;
        let mut char_count = 0;
        let mut last_char: Option<char> = None;

        fn write_compressed(chars: &mut [char], write_index: &mut usize, c: char, count: i32) {
            chars[*write_index] = c;
            *write_index += 1;

            if count == 1 {
                return;
            }

            for ch in count.to_string().chars() {
                chars[*write_index] = ch;
                *write_index += 1;
            }
        }

        for i in 0..chars.len() {
            let x = chars[i];

            if last_char.is_some() && last_char.unwrap() == x {
                char_count += 1;
                continue;
            }

            if let Some(c) = last_char {
                write_compressed(chars, &mut write_idx, c, char_count);
            }

            last_char = Some(x);
            char_count = 1;
        }

        if let Some(c) = last_char {
            write_compressed(chars, &mut write_idx, c, char_count);
        }

        write_idx as i32
    }
}
