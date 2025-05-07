struct Solution;

impl Solution {
    pub fn decode_string(s: String) -> String {
        let mut stack: Vec<(String, usize)> = Vec::new();
        let mut curr = String::new();
        let mut num = 0;

        for c in s.chars() {
            if c.is_ascii_digit() {
                num = num * 10 + c.to_digit(10).unwrap() as usize;
            } else if c == '[' {
                stack.push((curr.clone(), num));
                curr.clear();
                num = 0;
            } else if c == ']' {
                let (prev, times) = stack.pop().unwrap();
                curr = prev + &curr.repeat(times);
            } else {
                curr.push(c);
            }
        }

        curr
    }

    pub fn decode_string_old(s: String) -> String {
        let mut out = String::new();

        let chars: Vec<_> = s.chars().collect();

        let mut i = 0;
        while i < chars.len() {
            let c = chars[i];

            if c.is_alphabetic() {
                out.push(c);
                i += 1;
                continue;
            }

            if c.is_numeric() {
                let start = i;
                let mut end = i;
                while chars[end].is_numeric() {
                    end += 1;
                }

                let p = &s[start..end];
                let repeat_times = p.parse::<usize>().unwrap();

                let mut brackets = 0;

                let mut j = end;
                loop {
                    match chars[j] {
                        '[' => brackets += 1,
                        ']' => brackets -= 1,
                        _ => {}
                    }
                    if brackets == 0 {
                        break;
                    }
                    j += 1
                }

                let res = Self::decode_string(s[end + 1..j].to_string());
                out.push_str(res.repeat(repeat_times).as_str());

                i = j + 1
            }
        }

        out
    }
}

#[test]
fn test() {
    let test_cases = [
        ("3[a]2[bc]", "aaabcbc"),
        ("3[a2[c]]", "accaccacc"),
        ("2[abc]3[cd]ef", "abcabccdcdcdef"),
    ];

    for (input, expected) in test_cases {
        assert_eq!(
            Solution::decode_string(input.to_string()),
            expected,
            "failed for input '{}'",
            input
        );
        assert_eq!(
            Solution::decode_string_old(input.to_string()),
            expected,
            "failed for input '{}'",
            input
        );
    }
}
