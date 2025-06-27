use std::collections::HashSet;

struct Solution;

impl Solution {
    // https://leetcode.com/problems/reachable-nodes-with-restrictions/
    pub fn reachable_nodes(n: i32, edges: Vec<Vec<i32>>, restricted: Vec<i32>) -> i32 {
        let restricted_nodes: HashSet<i32> = restricted.into_iter().collect();

        let mut adj_list = vec![vec![]; n as usize];

        for edge in &edges {
            adj_list[edge[0] as usize].push(edge[1]);
            adj_list[edge[1] as usize].push(edge[0]);
        }

        let mut visited_count = 0;
        let mut visited = vec![false; n as usize];

        let mut stack = vec![0];
        while let Some(node) = stack.pop() {
            if visited[node as usize] || restricted_nodes.contains(&node) {
                continue;
            }
            visited[node as usize] = true;

            visited_count += 1;
            for &edge in adj_list[node as usize].iter() {
                stack.push(edge);
            }
        }

        visited_count
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            7,
            vec![[0, 1], [0, 2], [0, 5], [0, 4], [3, 2], [6, 5]],
            vec![4, 2, 1],
            3,
        ),
        (
            7,
            vec![[0, 1], [1, 2], [3, 1], [4, 0], [0, 5], [5, 6]],
            vec![4, 5],
            4,
        ),
    ];

    for (n, edges, restricted, expected) in test_cases {
        assert_eq!(
            Solution::reachable_nodes(n, edges.iter().map(Vec::from).collect(), restricted.clone()),
            expected,
            "failed for n = {n}, edges = {edges:?}, restricted = {restricted:?}",
        );
    }
}
