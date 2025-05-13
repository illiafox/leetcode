struct Solution;

impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        let mut num = x;
        let mut rev = 0;
        while num > 0 {
            let dig = num % 10;
            rev = rev * 10 + dig;
            num /= 10;
        }

        x == rev
    }
}

#[test]
fn test() {
    let test_cases = [(121, true), (-121, false), (10, false)];

    for (x, expected) in test_cases {
        assert_eq!(Solution::is_palindrome(x), expected, "failed for {x}",);
    }
}
