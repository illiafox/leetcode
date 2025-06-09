struct Solution;

impl Solution {
    fn traverse(edges: &[i32], start_node: usize) -> Vec<i32> {
        let mut distances = vec![-1; edges.len()];
        let mut current = start_node;
        let mut distance = 0;

        while distances[current] == -1 {
            distances[current] = distance;
            let idx = edges[current];
            if idx == -1 {
                break;
            }
            current = idx as usize;
            distance += 1;
        }

        distances
    }

    // https://leetcode.com/problems/find-closest-node-to-given-two-nodes
    // I am doing this again...
    pub fn closest_meeting_node(edges: Vec<i32>, node1: i32, node2: i32) -> i32 {
        let node1_distances = Self::traverse(&edges, node1 as usize);
        let node2_distances = Self::traverse(&edges, node2 as usize);

        let mut candidate: i32 = -1;
        let mut min_distance = i32::MAX;

        for i in 0..edges.len() {
            if node1_distances[i] == -1 || node2_distances[i] == -1 {
                continue;
            }

            let distance = node2_distances[i].max(node1_distances[i]);
            if distance < min_distance {
                candidate = i as i32;
                min_distance = distance;
            }
        }

        candidate
    }
}

#[test]
fn test() {
    let test_cases = vec![
        (vec![2, 2, 3, -1], 0, 1, 2),
        (vec![1, 2, -1], 0, 2, 2),
        (vec![4, 4, 4, 5, 1, 2, 2], 1, 1, 1),
        (vec![5, 3, 1, 0, 2, 4, 5], 3, 2, 3),
    ];

    for (edges, node1, node2, expected) in test_cases {
        assert_eq!(
            Solution::closest_meeting_node(edges.clone(), node1, node2),
            expected,
            "closest_meeting_node: failed for edges: {edges:?}, node1 {node1} node2 {node2}",
        );
    }
}
