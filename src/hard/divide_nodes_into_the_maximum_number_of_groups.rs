use std::cmp::max;
use std::collections::VecDeque;

struct Solution;

impl Solution {
    fn is_bipartite(adj_list: &[Vec<i32>], node: usize, colors: &mut Vec<i8>) -> bool {
        for &neighbor in &adj_list[node] {
            let neighbor = neighbor as usize;

            if colors[neighbor] == colors[node] {
                return false;
            }

            if colors[neighbor] != -1 {
                continue;
            }

            colors[neighbor] = (colors[node] + 1) % 2;

            if !Self::is_bipartite(adj_list, neighbor, colors) {
                return false;
            }
        }
        true
    }

    fn longest_shortest_path(adj_list: &[Vec<i32>], node: usize, n: usize) -> i32 {
        let mut queue = VecDeque::new();
        let mut visited = vec![false; n];

        queue.push_back(node);
        visited[node] = true;

        let mut distance = 0;

        while !queue.is_empty() {
            let current_queue_size = queue.len();

            for _ in 0..current_queue_size {
                let current_node = queue.pop_front().unwrap();

                for &neighbor in &adj_list[current_node] {
                    let neighbor = neighbor as usize;

                    if visited[neighbor] {
                        continue;
                    }

                    visited[neighbor] = true;
                    queue.push_back(neighbor);
                }
            }

            distance += 1;
        }

        distance
    }

    fn get_numbers_of_groups(
        edges: &Vec<Vec<i32>>,
        node: usize,
        distances: &Vec<i32>,
        visited: &mut Vec<bool>,
    ) -> i32 {
        let mut max_number_of_groups = distances[node];
        visited[node] = true;

        for &neighbor in &edges[node] {
            let neighbor = neighbor as usize;

            if visited[neighbor] {
                continue;
            }

            max_number_of_groups = max(
                max_number_of_groups,
                Self::get_numbers_of_groups(edges, neighbor, distances, visited),
            );
        }

        max_number_of_groups
    }

    pub fn magnificent_sets(n: i32, edges: Vec<Vec<i32>>) -> i32 {
        let n = n as usize;

        let mut adj_list: Vec<Vec<i32>> = vec![Vec::new(); n];

        for edge in edges {
            adj_list[(edge[0] - 1) as usize].push(edge[1] - 1);
            adj_list[(edge[1] - 1) as usize].push(edge[0] - 1);
        }

        let mut colors = vec![-1; n];

        for node in 0..n {
            if colors[node] != -1 {
                continue;
            }

            colors[node] = 0;

            if !Self::is_bipartite(&adj_list, node, &mut colors) {
                return -1;
            }
        }

        let mut distances = vec![0; n];

        for (node, distance) in distances.iter_mut().enumerate() {
            *distance = Self::longest_shortest_path(&adj_list, node, n)
        }

        let mut max_number_of_groups = 0;

        let mut visited = vec![false; n];

        for node in 0..n {
            if visited[node] {
                continue;
            }

            colors[node] = 0;

            max_number_of_groups +=
                Self::get_numbers_of_groups(&adj_list, node, &distances, &mut visited)
        }

        max_number_of_groups
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            6,
            vec![
                vec![1, 2],
                vec![1, 4],
                vec![1, 5],
                vec![2, 6],
                vec![2, 3],
                vec![4, 6],
            ],
            4,
        ),
        (3, vec![vec![1, 2], vec![2, 3], vec![3, 1]], -1),
    ];

    for (n, edges, expected) in test_cases {
        assert_eq!(
            Solution::magnificent_sets(n, edges.clone()),
            expected,
            "failed for input: n: {} edges: {:?}",
            n,
            edges
        );
    }
}
