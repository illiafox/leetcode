struct Solution;

impl Solution {
    pub fn judge_circle(moves: String) -> bool {
        let mut x = 0;
        let mut y = 0;

        for c in moves.chars() {
            match c {
                'R' => x += 1,
                'L' => x -= 1,
                'U' => y += 1,
                'D' => y -= 1,
                _ => unreachable!(),
            }
        }

        x == 0 && y == 0
    }
}

#[test]
fn test() {
    let test_cases = [("UD", true), ("LL", false)];

    for (moves, expected) in test_cases {
        assert_eq!(
            Solution::judge_circle(moves.to_string()),
            expected,
            "failed for moves {moves}",
        );
    }
}
