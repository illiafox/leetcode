struct Solution;

impl Solution {
    pub fn can_be_valid(s: String, locked: String) -> bool {
        if s.len() % 2 == 1 {
            // odd string cannot be valid
            return false;
        }

        let mut open_positions: Vec<usize> = vec![];
        let mut unlocked_positions: Vec<usize> = vec![];

        let locked_info = locked.as_bytes();

        for (i, &ch) in s.as_bytes().iter().enumerate() {
            match (locked_info[i], ch) {
                (b'0', _) => unlocked_positions.push(i), // parenthesis is unlocked
                (_, b'(') => open_positions.push(i),     // open parenthesis
                (_, b')') => {
                    // match open or unlocked parenthesis for closed one
                    if open_positions.pop().is_none() && unlocked_positions.pop().is_none() {
                        // Neither are available
                        return false;
                    }
                }
                _ => {}
            }
        }

        // match open parentheses left with available unlocked ones
        while let (Some(open_idx), Some(unlocked_idx)) =
            (open_positions.pop(), unlocked_positions.pop())
        {
            if unlocked_idx < open_idx {
                return false;
            }
        }

        open_positions.is_empty()
    }
}

#[test]
fn test() {
    let test_cases = [
        ("))()))", "010100", true),
        ("()()", "0000", true),
        (")", "0", false),
    ];

    for &(s, locked, expected) in &test_cases {
        assert_eq!(
            Solution::can_be_valid(s.to_string(), locked.to_string()),
            expected,
            "failed for input s: {}, locked: {}",
            s,
            locked
        );
    }
}
