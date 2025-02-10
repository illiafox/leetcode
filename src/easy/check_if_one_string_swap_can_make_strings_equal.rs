struct Solution;

impl Solution {
    pub fn are_almost_equal(s1: String, s2: String) -> bool {
        if s1.len() != s2.len() {
            return false;
        }

        let (s1_bytes, s2_bytes) = (s1.as_bytes(), s2.as_bytes());

        let mut differences = 0;
        let mut difs: [(u8, u8); 2] = [(0, 0), (0, 0)];

        for (b1, b2) in s1_bytes.iter().zip(s2_bytes.iter()) {
            if b1 != b2 {
                if differences == 2 {
                    return false;
                }
                difs[differences] = (*b1, *b2);
                differences += 1;
            }
        }

        differences == 0 || (differences == 2 && difs[0].0 == difs[1].1 && difs[0].1 == difs[1].0)
    }
}

#[test]
fn test() {
    let test_cases = [
        ("bank", "kanb", true),
        ("attack", "defend", false),
        ("kelb", "kelb", true),
        ("caa", "aaz", false),
    ];

    for &(s1, s2, expected) in &test_cases {
        assert_eq!(
            Solution::are_almost_equal(s1.into(), s2.into()),
            expected,
            "failed for s1: {}, s2: {} k: {}",
            s1,
            s2,
            expected
        );
    }
}
