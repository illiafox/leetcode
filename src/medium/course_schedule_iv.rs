use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn is_prerequisite(
        visited_courses: &mut Vec<bool>,
        adj_list: &HashMap<i32, Vec<i32>>,
        src: i32,
        target: i32,
    ) -> bool {
        visited_courses[src as usize] = true;

        if src == target {
            return true;
        }

        if let Some(neighbors) = adj_list.get(&src) {
            for &neighbor in neighbors {
                if !visited_courses[neighbor as usize]
                    && Self::is_prerequisite(visited_courses, adj_list, neighbor, target)
                {
                    return true;
                }
            }
        }

        false
    }

    pub fn check_if_prerequisite(
        num_courses: i32,
        prerequisites: Vec<Vec<i32>>,
        queries: Vec<Vec<i32>>,
    ) -> Vec<bool> {
        let mut adj_list: HashMap<i32, Vec<i32>> = HashMap::new();

        for node in prerequisites {
            adj_list.entry(node[0]).or_default().push(node[1]);
        }

        let mut visited_courses = vec![false; num_courses as usize];

        queries
            .iter()
            .map(|x| {
                visited_courses.fill(false);
                Self::is_prerequisite(&mut visited_courses, &adj_list, x[0], x[1])
            })
            .collect()
    }
}

#[test]
fn test() {
    let test_cases = [(
        2,
        vec![vec![1, 0]],
        vec![vec![0, 1], vec![1, 0]],
        vec![false, true],
    )];

    for (num_courses, prerequisites, queries, expected) in test_cases {
        assert_eq!(
            Solution::check_if_prerequisite(num_courses, prerequisites.clone(), queries.clone()),
            expected,
            "failed for input: num_courses: {} prerequisites: {:?} queries: {:?}",
            num_courses,
            prerequisites,
            queries
        );
    }
}
