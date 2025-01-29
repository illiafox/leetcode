struct Dsu {
    size: Vec<usize>,
    representative: Vec<usize>,
}

impl Dsu {
    pub fn new(n: usize) -> Self {
        let mut representative = vec![0; n];
        for (i, slot) in representative.iter_mut().enumerate() {
            *slot = i;
        }

        Dsu {
            size: vec![1; n],
            representative,
        }
    }

    pub fn find(&mut self, node: usize) -> usize {
        if self.representative[node] == node {
            return node;
        }

        let res = self.find(self.representative[node]);

        // path compression, next find will be faster
        self.representative[node] = res;

        res
    }

    pub fn do_union(&mut self, mut node_one: usize, mut node_two: usize) -> bool {
        node_one = self.find(node_one);
        node_two = self.find(node_two);

        if node_one == node_two {
            return false;

            // union by size, keep the tree as shallow as possible
        } else if self.size[node_one] > self.size[node_two] {
            self.representative[node_two] = node_one;
            self.size[node_one] += self.size[node_two];
        } else {
            self.representative[node_one] = node_two;
            self.size[node_two] += self.size[node_one];
        }

        true
    }
}

struct Solution;

impl Solution {
    pub fn find_redundant_connection(edges: Vec<Vec<i32>>) -> Vec<i32> {
        let n = edges.len();

        let mut d = Dsu::new(n);
        for edge in edges {
            if !d.do_union((edge[0] - 1) as usize, (edge[1] - 1) as usize) {
                return edge;
            }
        }

        vec![]
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![vec![1, 2], vec![1, 3], vec![2, 3]], vec![2, 3]),
        (
            vec![vec![1, 2], vec![2, 3], vec![3, 4], vec![1, 4], vec![1, 5]],
            vec![1, 4],
        ),
    ];

    for (edges, expected) in test_cases {
        assert_eq!(
            Solution::find_redundant_connection(edges.clone()),
            expected,
            "failed for input edges: {:?}",
            edges
        );
    }
}
