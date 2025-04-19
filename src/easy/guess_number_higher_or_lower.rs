struct Solution;

fn guess(_: i32) -> i32 {
    panic!("implement me")
}

impl Solution {
    unsafe fn guess_number(n: i32) -> i32 {
        let mut a = 1;
        let mut b = n + 1;
        let mut mid = n;

        loop {
            match guess(mid) {
                -1 => b = mid,
                1 => a = mid,
                0 => return mid,
                _ => panic!("what"),
            }

            mid = ((a as i64 + b as i64) / 2) as i32;
        }
    }
}
