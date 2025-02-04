use std::cmp::Ordering;
use std::collections::BinaryHeap;

#[derive(Debug, Copy, Clone)]
struct Edge {
    to: usize,
    cost: f64,
}

#[derive(Debug, PartialEq)]
struct State {
    index: usize,
    cost: f64,
}

impl Eq for State {}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        // flip ordering on `cost` so the smallest cost has higher priority in the heap
        other.cost.partial_cmp(&self.cost).unwrap()
    }
}

struct Solution;

impl Solution {
    pub fn max_probability(
        n: i32,
        edges: Vec<Vec<i32>>,
        succ_prob: Vec<f64>,
        start_node: i32,
        end_node: i32,
    ) -> f64 {
        let n = n as usize;
        let start_node = start_node as usize;
        let end_node = end_node as usize;

        // build adjacency list with "cost" = -ln(probability)
        let mut adj_list: Vec<Vec<Edge>> = vec![Vec::new(); n];
        for (i, edge) in edges.iter().enumerate() {
            let (u, v) = (edge[0] as usize, edge[1] as usize);
            let prob = succ_prob[i];

            // avoid taking ln(0.0). if prob can be zero, handle it as "unreachable"
            if prob > 0.0 {
                let cost = -prob.ln();
                adj_list[u].push(Edge { to: v, cost });
                adj_list[v].push(Edge { to: u, cost });
            }
        }

        let mut dist = vec![f64::INFINITY; n];
        dist[start_node] = 0.0;

        let mut heap = BinaryHeap::new();
        heap.push(State {
            index: start_node,
            cost: 0.0,
        });

        while let Some(State { index, cost }) = heap.pop() {
            // if we've found a better route already, skip
            if cost > dist[index] {
                continue;
            }

            if index == end_node {
                // cost = sum of (-ln(probabilities)) = -ln(product)
                // so product = exp(-cost)
                return (-cost).exp();
            }

            for edge in &adj_list[index] {
                let next_cost = cost + edge.cost;

                if next_cost < dist[edge.to] {
                    dist[edge.to] = next_cost;
                    heap.push(State {
                        index: edge.to,
                        cost: next_cost,
                    });
                }
            }
        }

        0.0
    }
}

#[test]
fn test() {
    let test_cases = [
        (
            3,
            vec![vec![0, 1], vec![1, 2], vec![0, 2]],
            vec![0.5, 0.5, 0.2],
            0,
            2,
            0.25000,
        ),
        (
            5,
            vec![
                vec![1, 4],
                vec![2, 4],
                vec![0, 4],
                vec![0, 3],
                vec![0, 2],
                vec![2, 3],
            ],
            vec![0.37, 0.17, 0.93, 0.23, 0.39, 0.04],
            3,
            4,
            0.2139,
        ),
    ];

    for (n, edges, succ_rates, start, end, expected) in test_cases {
        // "answer will be accepted if it differs from the correct answer by at most 1e-5"
        let eps = 1e-5;

        let result = Solution::max_probability(n, edges.clone(), succ_rates.clone(), start, end);
        assert!(
            (result - expected).abs() < eps,
            "failed for input: n: {} edges: {:?}; {result} != {expected}",
            n,
            edges
        );
    }
}
