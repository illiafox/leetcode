use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn query_results(_limit: i32, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut current_colors: HashMap<i32, i32> = HashMap::new(); // ball -> color

        let mut color_count: HashMap<i32, i32> = HashMap::new(); // color -> balls count

        let mut result = vec![0; queries.len()];

        for (i, query) in queries.iter().enumerate() {
            let ball = query[0];

            if let Some(c) = current_colors.get(&ball) {
                let s = color_count.get_mut(c).unwrap();
                *s -= 1;
                if *s == 0 {
                    color_count.remove(c);
                }
            }

            let new_color = query[1];
            *color_count.entry(new_color).or_default() += 1;
            current_colors.insert(ball, new_color);

            result[i] = color_count.len() as i32;
        }

        result
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            4,
            vec![vec![1, 4], vec![2, 5], vec![1, 3], vec![3, 4]],
            vec![1, 2, 2, 3],
        ),
        (
            4,
            vec![vec![0, 1], vec![1, 2], vec![2, 2], vec![3, 4], vec![4, 5]],
            vec![1, 2, 2, 3, 4],
        ),
    ];

    for (limit, queries, expected) in test_cases {
        assert_eq!(
            Solution::query_results(limit, queries.clone()),
            expected,
            "failed for input limit: {} grid: {:?}",
            limit,
            queries
        );
    }
}
