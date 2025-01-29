use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn lexicographically_smallest_array(nums: Vec<i32>, limit: i32) -> Vec<i32> {
        let mut nums_sorted = nums.clone();
        nums_sorted.sort_unstable();

        let mut num_to_group: HashMap<i32, usize> = HashMap::new();
        let mut groups: Vec<Vec<i32>> = Vec::new();

        let mut current_group_index = 0;

        groups.push(vec![nums_sorted[0]]);
        num_to_group.insert(nums_sorted[0], current_group_index);

        for (idx, &num) in nums_sorted.iter().enumerate() {
            if idx == 0 {
                continue;
            }

            if num - *groups[current_group_index].last().unwrap() <= limit {
                groups[current_group_index].push(num);
            } else {
                current_group_index += 1;
                groups.push(vec![num])
            }

            num_to_group.insert(num, current_group_index);
        }

        let mut nums_modified = nums.clone();

        let mut group_used = vec![0usize; groups.len()];

        for val in nums_modified.iter_mut() {
            let group_idx = *num_to_group.get(val).unwrap();

            *val = groups[group_idx][group_used[group_idx]];
            group_used[group_idx] += 1
        }

        nums_modified
    }
}

#[test]
fn test() {
    let test_cases = [
        (vec![1, 5, 3, 9, 8], 2, vec![1, 3, 5, 8, 9]),
        (vec![1, 7, 6, 18, 2, 1], 3, vec![1, 6, 7, 18, 1, 2]),
        (vec![1, 7, 28, 19, 10], 3, vec![1, 7, 28, 19, 10]),
        (
            vec![4, 3, 23, 84, 34, 88, 44, 44, 18, 15],
            3,
            vec![3, 4, 23, 84, 34, 88, 44, 44, 15, 18],
        ),
    ];

    for (nums, limit, expected) in test_cases {
        assert_eq!(
            Solution::lexicographically_smallest_array(nums.clone(), limit),
            expected,
            "failed for nums: {:?}",
            nums,
        );
    }
}
