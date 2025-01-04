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

#[test]
fn test() {
    let chars = &mut vec![
        'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b',
    ];
    let expected = 4;
    assert_eq!(Solution::compress(chars), expected);
    assert_eq!(chars[..expected as usize], vec!['a', 'b', '1', '2']);

    let chars = &mut vec!['a'];
    let expected = 1;
    assert_eq!(Solution::compress(chars), expected);
    assert_eq!(chars[..expected as usize], vec!['a']);

    let chars = &mut vec!['a', 'a', 'b', 'b', 'c', 'c', 'c'];
    let expected = 6;
    assert_eq!(Solution::compress(chars), expected);
    assert_eq!(
        chars[..expected as usize],
        vec!['a', '2', 'b', '2', 'c', '3']
    );
}
