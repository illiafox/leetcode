
struct Solution;

impl Solution {
    // https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box
    pub fn min_operations(boxes: String) -> Vec<i32> {
        let b = boxes.as_bytes();

        let mut total_balls_left = 0;
        let mut total_moves_left = 0;

        let mut total_balls_right = 0;
        let mut total_moves_right = 0;

        let mut answer: Vec<i32> = vec![0; b.len()];

        for i in 0..b.len() {
            answer[i] += total_moves_left;
            total_balls_left += (b[i] - b'0') as i32;
            total_moves_left += total_balls_left;

            let j = boxes.len() - i - 1;
            answer[j] += total_moves_right;
            total_balls_right += (b[j] - b'0') as i32;
            total_moves_right += total_balls_right;
        }

        answer
    }
}

#[test]
fn test() {
    assert_eq!(Solution::min_operations("110".to_string()), vec![1, 1, 3]);
    assert_eq!(
        Solution::min_operations("001011".to_string()),
        vec![11, 8, 5, 4, 3, 4]
    );
}
