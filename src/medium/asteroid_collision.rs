struct Solution;

use std::cmp::Ordering::{Equal, Less};

impl Solution {
    pub fn asteroid_collision(asteroids: Vec<i32>) -> Vec<i32> {
        let mut stack: Vec<i32> = Vec::new();

        for &ast in &asteroids {
            let mut current = ast;

            while let Some(&last) = stack.last() {
                if last > 0 && current < 0 {
                    match last.cmp(&-current) {
                        Less => {
                            stack.pop();
                            continue;
                        }
                        Equal => {
                            stack.pop();
                        }
                        _ => {}
                    }

                    current = 0;
                    break;
                } else {
                    break;
                }
            }

            if current != 0 {
                stack.push(current)
            }
        }

        stack
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![-2, -1, 1, 2], vec![-2, -1, 1, 2]),
        (vec![5, 10, -5], vec![5, 10]),
        (vec![8, -8], vec![]),
        (vec![10, 2, -5], vec![10]),
        (vec![-2, -1, 1, 2], vec![-2, -1, 1, 2]),
    ];

    for (input, expected) in test_cases {
        assert_eq!(
            Solution::asteroid_collision(input.clone()),
            expected,
            "failed for input: {:?}",
            input.clone(),
        );
    }
}
