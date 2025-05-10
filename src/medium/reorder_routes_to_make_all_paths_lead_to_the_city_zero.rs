struct Solution;

use std::collections::{HashMap, HashSet, VecDeque};

impl Solution {
    pub fn min_reorder_v2(n: i32, connections: Vec<Vec<i32>>) -> i32 {
        let mut adj = vec![vec![]; n as usize];

        for edge in &connections {
            let (u, v) = (edge[0] as usize, edge[1] as usize);
            // store both directions
            // true means original direction u → v
            adj[u].push((v, true));
            adj[v].push((u, false)); // reverse direction
        }

        let mut visited = vec![false; n as usize];
        let mut queue = VecDeque::new();
        let mut swaps = 0;

        visited[0] = true;
        queue.push_back(0);

        while let Some(node) = queue.pop_front() {
            for &(neighbor, needs_reversal) in &adj[node] {
                if !visited[neighbor] {
                    visited[neighbor] = true;
                    if needs_reversal {
                        swaps += 1;
                    }
                    queue.push_back(neighbor);
                }
            }
        }

        swaps
    }
    pub fn min_reorder(n: i32, connections: Vec<Vec<i32>>) -> i32 {
        let mut adj_list: HashMap<i32, Vec<i32>> = HashMap::new();
        let mut directed_edges: HashSet<(i32, i32)> = HashSet::new();

        for edge in connections {
            let (u, v) = (edge[0], edge[1]);
            adj_list.entry(u).or_default().push(v);
            adj_list.entry(v).or_default().push(u);
            directed_edges.insert((u, v));
        }

        let mut visited = vec![false; n as usize];
        let mut swaps = 0;

        let mut queue = VecDeque::new();
        queue.push_back(0);
        visited[0] = true;

        while let Some(current) = queue.pop_front() {
            if let Some(neighbors) = adj_list.get(&current) {
                for &neighbor in neighbors {
                    if !visited[neighbor as usize] {
                        // reversed if the original direction is from current to neighbor
                        if directed_edges.contains(&(current, neighbor)) {
                            swaps += 1;
                        }
                        visited[neighbor as usize] = true;
                        queue.push_back(neighbor);
                    }
                }
            }
        }

        swaps
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            6,
            vec![vec![0, 1], vec![1, 3], vec![2, 3], vec![4, 0], vec![4, 5]],
            3,
        ),
        (5, vec![vec![1, 0], vec![1, 2], vec![3, 2], vec![3, 4]], 2),
        (6, vec![vec![1, 0], vec![2, 0]], 0),
    ];

    for (n, connections, expected) in test_cases {
        assert_eq!(
            Solution::min_reorder(n, connections.clone()),
            expected,
            "failed for n {} connections: {:?}",
            n,
            connections
        );
        assert_eq!(
            Solution::min_reorder_v2(n, connections.clone()),
            expected,
            "failed for n {} connections: {:?}",
            n,
            connections
        );
    }
}
