use std::collections::VecDeque;

struct Solution;

impl Solution {
    pub fn eventual_safe_nodes(graph: Vec<Vec<i32>>) -> Vec<i32> {
        let n = graph.len();
        let mut indegree = vec![0; n];
        let mut adj = vec![vec![]; n];

        // Build the reversed graph and calculate the indegree for each node
        for (i, neighbors) in graph.iter().enumerate() {
            for &node in neighbors {
                adj[node as usize].push(i);
                indegree[i] += 1;
            }
        }

        let mut queue = VecDeque::new();
        // Push all nodes with indegree 0 into the queue
        for (i, &indegree) in indegree.iter().enumerate() {
            if indegree == 0 {
                queue.push_back(i);
            }
        }

        let mut safe = vec![false; n];
        while let Some(node) = queue.pop_front() {
            safe[node] = true;

            for &neighbor in &adj[node] {
                // Decrease the indegree of the neighbor
                indegree[neighbor] -= 1;
                if indegree[neighbor] == 0 {
                    queue.push_back(neighbor);
                }
            }
        }

        // Collect all the safe nodes
        let mut safe_nodes = vec![];
        for (i, &is_safe) in safe.iter().enumerate() {
            if is_safe {
                safe_nodes.push(i as i32);
            }
        }

        safe_nodes
    }
}

#[test]
fn test() {
    let test_cases = vec![
        // Test case 1
        (
            vec![
                vec![1, 2],
                vec![2, 3],
                vec![5],
                vec![0],
                vec![5],
                vec![],
                vec![],
            ],
            vec![2, 4, 5, 6],
        ),
        // Test case 2
        (
            vec![vec![1, 2, 3, 4], vec![1, 2], vec![3, 4], vec![0, 4], vec![]],
            vec![4],
        ),
        // Test case 3
        (
            vec![vec![], vec![0, 2, 3, 4], vec![3], vec![4], vec![]],
            vec![0, 1, 2, 3, 4],
        ),
    ];

    for (graph, expected) in test_cases {
        assert_eq!(
            Solution::eventual_safe_nodes(graph.clone()),
            expected,
            "eventual_safe_nodes: failed for graph: {:?}",
            graph
        );
    }
}
